import {library} from '@fortawesome/fontawesome-svg-core'
import {
	faAlignLeft,
	faAngleRight,
	faAnglesUp,
	faArchive,
	faArrowLeft,
	faArrowUpFromBracket,
	faBold,
	faItalic,
	faStrikethrough,
	faCode,
	faBars,
	faBell,
	faBolt,
	faCalendar,
	faCheck,
	faCheckDouble,
	faChessKnight,
	faChevronDown,
	faCircleInfo,
	faCloudDownloadAlt,
	faCloudUploadAlt,
	faCocktail,
	faCoffee,
	faCog,
	faCopy,
	faDownload,
	faEllipsisH,
	faEllipsisV,
	faExclamationCircle,
	faEye,
	faEyeSlash,
	faFile,
	faFileImage,
	faFillDrip,
	faFilter,
	faForward,
	faGripLines,
	faHeader,
	faHistory,
	faImage,
	faKeyboard,
	faLayerGroup,
	faList,
	faListOl,
	faLock,
	faPaperclip,
	faPaste,
	faPen,
	faPencilAlt,
	faPercent,
	faPlay,
	faPlus,
	faPowerOff,
	faSearch,
	faShareAlt,
	faSignOutAlt,
	faSitemap,
	faSort,
	faSortUp,
	faStar as faStarSolid,
	faStop,
	faTachometerAlt,
	faTags,
	faTasks,
	faTh,
	faTimes,
	faTrashAlt,
	faUser,
	faUsers,
	faQuoteRight,
	faListUl,
	faLink,
	faUndo,
	faRedo,
	faUnlink,
	faParagraph,
	faTable,
	faX, faArrowTurnDown, faListCheck, faXmark, faXmarksLines, faFont, faRulerHorizontal, faUnderline,
} from '@fortawesome/free-solid-svg-icons'
import {
	faBellSlash,
	faCalendarAlt,
	faCheckSquare,
	faClock,
	faComments,
	faSave,
	faSquareCheck,
	faStar,
	faSun,
	faTimesCircle,
	faCircleQuestion, faFaceLaugh,
} from '@fortawesome/free-regular-svg-icons'
import {FontAwesomeIcon} from '@fortawesome/vue-fontawesome'

import type {FontAwesomeIcon as FontAwesomeIconFixedTypes} from '@/types/vue-fontawesome'

library.add(faBold)
library.add(faUndo)
library.add(faRedo)
library.add(faItalic)
library.add(faLink)
library.add(faUnlink)
library.add(faParagraph)
library.add(faSquareCheck)
library.add(faTable)
library.add(faFile)
library.add(faFileImage)
library.add(faCheckSquare)
library.add(faStrikethrough)
library.add(faCode)
library.add(faQuoteRight)
library.add(faListUl)
library.add(faAlignLeft)
library.add(faAngleRight)
library.add(faArchive)
library.add(faArrowLeft)
library.add(faBars)
library.add(faBell)
library.add(faBellSlash)
library.add(faCalendar)
library.add(faCalendarAlt)
library.add(faCheck)
library.add(faCheckDouble)
library.add(faChessKnight)
library.add(faChevronDown)
library.add(faCircleInfo)
library.add(faCircleQuestion)
library.add(faClock)
library.add(faCloudDownloadAlt)
library.add(faCloudUploadAlt)
library.add(faCocktail)
library.add(faCoffee)
library.add(faCog)
library.add(faComments)
library.add(faCopy)
library.add(faDownload)
library.add(faEllipsisH)
library.add(faEllipsisV)
library.add(faExclamationCircle)
library.add(faEye)
library.add(faEyeSlash)
library.add(faFillDrip)
library.add(faFilter)
library.add(faForward)
library.add(faGripLines)
library.add(faHeader)
library.add(faHistory)
library.add(faImage)
library.add(faKeyboard)
library.add(faLayerGroup)
library.add(faList)
library.add(faListOl)
library.add(faLock)
library.add(faPaperclip)
library.add(faPaste)
library.add(faPen)
library.add(faPencilAlt)
library.add(faPercent)
library.add(faPlay)
library.add(faPlus)
library.add(faPowerOff)
library.add(faSave)
library.add(faSearch)
library.add(faShareAlt)
library.add(faSignOutAlt)
library.add(faSitemap)
library.add(faSort)
library.add(faSortUp)
library.add(faStar)
library.add(faStarSolid)
library.add(faStop)
library.add(faSun)
library.add(faTachometerAlt)
library.add(faTags)
library.add(faTasks)
library.add(faTh)
library.add(faTimes)
library.add(faTimesCircle)
library.add(faTrashAlt)
library.add(faUser)
library.add(faUsers)
library.add(faArrowUpFromBracket)
library.add(faX)
library.add(faAnglesUp)
library.add(faBolt)
library.add(faArrowTurnDown)
library.add(faListCheck)
library.add(faXmark)
library.add(faXmarksLines)
library.add(faFont)
library.add(faRulerHorizontal)
library.add(faUnderline)
library.add(faFaceLaugh)

// overwriting the wrong types
export default FontAwesomeIcon as unknown as FontAwesomeIconFixedTypes