// Vikunja is a to-do list application to facilitate your life.
// Copyright 2018-2021 Vikunja and contributors. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public Licensee as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public Licensee for more details.
//
// You should have received a copy of the GNU Affero General Public Licensee
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package ticktick

import (
	"encoding/csv"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"

	"code.vikunja.io/api/pkg/models"
	"code.vikunja.io/api/pkg/modules/migration"
	"code.vikunja.io/api/pkg/user"
)

const timeISO = "2006-01-02T15:04:05-0700"

type Migrator struct {
}

type tickTickTask struct {
	FolderName    string
	ListName      string
	Title         string
	Tags          []string
	Content       string
	IsChecklist   bool
	StartDate     time.Time
	DueDate       time.Time
	Reminder      time.Duration
	Repeat        string
	Priority      int
	Status        string
	CreatedTime   time.Time
	CompletedTime time.Time
	Order         float64
	TaskID        int64
	ParentID      int64
}

func convertTickTickToVikunja(tasks []*tickTickTask) (result []*models.NamespaceWithListsAndTasks) {
	namespace := &models.NamespaceWithListsAndTasks{
		Namespace: models.Namespace{
			Title: "Migrated from TickTick",
		},
		Lists: []*models.ListWithTasksAndBuckets{},
	}

	lists := make(map[string]*models.ListWithTasksAndBuckets)
	for _, t := range tasks {
		_, has := lists[t.ListName]
		if !has {
			lists[t.ListName] = &models.ListWithTasksAndBuckets{
				List: models.List{
					Title: t.ListName,
				},
			}
		}

		labels := make([]*models.Label, 0, len(t.Tags))
		for _, tag := range t.Tags {
			labels = append(labels, &models.Label{
				Title: tag,
			})
		}

		task := &models.TaskWithComments{
			Task: models.Task{
				ID:          t.TaskID,
				Title:       t.Title,
				Description: t.Content,
				StartDate:   t.StartDate,
				EndDate:     t.DueDate,
				DueDate:     t.DueDate,
				Reminders: []time.Time{
					t.DueDate.Add(t.Reminder * -1),
				},
				Done:     t.Status == "1",
				DoneAt:   t.CompletedTime,
				Position: t.Order,
				Labels:   labels,
			},
		}

		if t.ParentID != 0 {
			task.RelatedTasks = map[models.RelationKind][]*models.Task{
				models.RelationKindParenttask: {{ID: t.ParentID}},
			}
		}

		lists[t.ListName].Tasks = append(lists[t.ListName].Tasks, task)
	}

	for _, l := range lists {
		namespace.Lists = append(namespace.Lists, l)
	}

	sort.Slice(namespace.Lists, func(i, j int) bool {
		return namespace.Lists[i].Title < namespace.Lists[j].Title
	})

	return []*models.NamespaceWithListsAndTasks{namespace}
}

// Name is used to get the name of the ticktick migration - we're using the docs here to annotate the status route.
// @Summary Get migration status
// @Description Returns if the current user already did the migation or not. This is useful to show a confirmation message in the frontend if the user is trying to do the same migration again.
// @tags migration
// @Produce json
// @Security JWTKeyAuth
// @Success 200 {object} migration.Status "The migration status"
// @Failure 500 {object} models.Message "Internal server error"
// @Router /migration/ticktick/status [get]
func (m *Migrator) Name() string {
	return "ticktick"
}

// Migrate takes a ticktick export, parses it and imports everything in it into Vikunja.
// @Summary Import all lists, tasks etc. from a TickTick backup export
// @Description Imports all projects, tasks, notes, reminders, subtasks and files from a TickTick backup export into Vikunja.
// @tags migration
// @Accept json
// @Produce json
// @Security JWTKeyAuth
// @Param import formData string true "The TickTick backup csv file."
// @Success 200 {object} models.Message "A message telling you everything was migrated successfully."
// @Failure 500 {object} models.Message "Internal server error"
// @Router /migration/ticktick/migrate [post]
func (m *Migrator) Migrate(user *user.User, file io.ReaderAt, size int64) error {
	fr := io.NewSectionReader(file, 0, 0)
	r := csv.NewReader(fr)
	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	allTasks := make([]*tickTickTask, 0, len(records))
	for line, record := range records {
		if line <= 3 {
			continue
		}
		startDate, err := time.Parse(timeISO, record[6])
		if err != nil {
			return err
		}
		dueDate, err := time.Parse(timeISO, record[7])
		if err != nil {
			return err
		}
		// TODO: parse properly
		reminder, err := time.ParseDuration(record[8])
		if err != nil {
			return err
		}
		priority, err := strconv.Atoi(record[10])
		if err != nil {
			return err
		}
		createdTime, err := time.Parse(timeISO, record[12])
		if err != nil {
			return err
		}
		completedTime, err := time.Parse(timeISO, record[13])
		if err != nil {
			return err
		}
		order, err := strconv.ParseFloat(record[14], 64)
		if err != nil {
			return err
		}
		taskID, err := strconv.ParseInt(record[21], 10, 64)
		if err != nil {
			return err
		}
		parentID, err := strconv.ParseInt(record[21], 10, 64)
		if err != nil {
			return err
		}

		allTasks = append(allTasks, &tickTickTask{
			ListName:      record[1],
			Title:         record[2],
			Tags:          strings.Split(record[3], ", "),
			Content:       record[4],
			IsChecklist:   record[5] == "Y",
			StartDate:     startDate,
			DueDate:       dueDate,
			Reminder:      reminder,
			Repeat:        record[9],
			Priority:      priority,
			Status:        record[11],
			CreatedTime:   createdTime,
			CompletedTime: completedTime,
			Order:         order,
			TaskID:        taskID,
			ParentID:      parentID,
		})
	}

	vikunjaTasks := convertTickTickToVikunja(allTasks)

	return migration.InsertFromStructure(vikunjaTasks, user)
}
