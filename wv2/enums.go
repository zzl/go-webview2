package wv2

// enum COREWEBVIEW2_KEY_EVENT_KIND
var COREWEBVIEW2_KEY_EVENT_KIND = struct {
	COREWEBVIEW2_KEY_EVENT_KIND_KEY_DOWN int32
	COREWEBVIEW2_KEY_EVENT_KIND_KEY_UP int32
	COREWEBVIEW2_KEY_EVENT_KIND_SYSTEM_KEY_DOWN int32
	COREWEBVIEW2_KEY_EVENT_KIND_SYSTEM_KEY_UP int32
}{
	COREWEBVIEW2_KEY_EVENT_KIND_KEY_DOWN: 0,
	COREWEBVIEW2_KEY_EVENT_KIND_KEY_UP: 1,
	COREWEBVIEW2_KEY_EVENT_KIND_SYSTEM_KEY_DOWN: 2,
	COREWEBVIEW2_KEY_EVENT_KIND_SYSTEM_KEY_UP: 3,
}

// enum COREWEBVIEW2_MOVE_FOCUS_REASON
var COREWEBVIEW2_MOVE_FOCUS_REASON = struct {
	COREWEBVIEW2_MOVE_FOCUS_REASON_PROGRAMMATIC int32
	COREWEBVIEW2_MOVE_FOCUS_REASON_NEXT int32
	COREWEBVIEW2_MOVE_FOCUS_REASON_PREVIOUS int32
}{
	COREWEBVIEW2_MOVE_FOCUS_REASON_PROGRAMMATIC: 0,
	COREWEBVIEW2_MOVE_FOCUS_REASON_NEXT: 1,
	COREWEBVIEW2_MOVE_FOCUS_REASON_PREVIOUS: 2,
}

// enum COREWEBVIEW2_WEB_ERROR_STATUS
var COREWEBVIEW2_WEB_ERROR_STATUS = struct {
	COREWEBVIEW2_WEB_ERROR_STATUS_UNKNOWN int32
	COREWEBVIEW2_WEB_ERROR_STATUS_CERTIFICATE_COMMON_NAME_IS_INCORRECT int32
	COREWEBVIEW2_WEB_ERROR_STATUS_CERTIFICATE_EXPIRED int32
	COREWEBVIEW2_WEB_ERROR_STATUS_CLIENT_CERTIFICATE_CONTAINS_ERRORS int32
	COREWEBVIEW2_WEB_ERROR_STATUS_CERTIFICATE_REVOKED int32
	COREWEBVIEW2_WEB_ERROR_STATUS_CERTIFICATE_IS_INVALID int32
	COREWEBVIEW2_WEB_ERROR_STATUS_SERVER_UNREACHABLE int32
	COREWEBVIEW2_WEB_ERROR_STATUS_TIMEOUT int32
	COREWEBVIEW2_WEB_ERROR_STATUS_ERROR_HTTP_INVALID_SERVER_RESPONSE int32
	COREWEBVIEW2_WEB_ERROR_STATUS_CONNECTION_ABORTED int32
	COREWEBVIEW2_WEB_ERROR_STATUS_CONNECTION_RESET int32
	COREWEBVIEW2_WEB_ERROR_STATUS_DISCONNECTED int32
	COREWEBVIEW2_WEB_ERROR_STATUS_CANNOT_CONNECT int32
	COREWEBVIEW2_WEB_ERROR_STATUS_HOST_NAME_NOT_RESOLVED int32
	COREWEBVIEW2_WEB_ERROR_STATUS_OPERATION_CANCELED int32
	COREWEBVIEW2_WEB_ERROR_STATUS_REDIRECT_FAILED int32
	COREWEBVIEW2_WEB_ERROR_STATUS_UNEXPECTED_ERROR int32
	COREWEBVIEW2_WEB_ERROR_STATUS_VALID_AUTHENTICATION_CREDENTIALS_REQUIRED int32
	COREWEBVIEW2_WEB_ERROR_STATUS_VALID_PROXY_AUTHENTICATION_REQUIRED int32
}{
	COREWEBVIEW2_WEB_ERROR_STATUS_UNKNOWN: 0,
	COREWEBVIEW2_WEB_ERROR_STATUS_CERTIFICATE_COMMON_NAME_IS_INCORRECT: 1,
	COREWEBVIEW2_WEB_ERROR_STATUS_CERTIFICATE_EXPIRED: 2,
	COREWEBVIEW2_WEB_ERROR_STATUS_CLIENT_CERTIFICATE_CONTAINS_ERRORS: 3,
	COREWEBVIEW2_WEB_ERROR_STATUS_CERTIFICATE_REVOKED: 4,
	COREWEBVIEW2_WEB_ERROR_STATUS_CERTIFICATE_IS_INVALID: 5,
	COREWEBVIEW2_WEB_ERROR_STATUS_SERVER_UNREACHABLE: 6,
	COREWEBVIEW2_WEB_ERROR_STATUS_TIMEOUT: 7,
	COREWEBVIEW2_WEB_ERROR_STATUS_ERROR_HTTP_INVALID_SERVER_RESPONSE: 8,
	COREWEBVIEW2_WEB_ERROR_STATUS_CONNECTION_ABORTED: 9,
	COREWEBVIEW2_WEB_ERROR_STATUS_CONNECTION_RESET: 10,
	COREWEBVIEW2_WEB_ERROR_STATUS_DISCONNECTED: 11,
	COREWEBVIEW2_WEB_ERROR_STATUS_CANNOT_CONNECT: 12,
	COREWEBVIEW2_WEB_ERROR_STATUS_HOST_NAME_NOT_RESOLVED: 13,
	COREWEBVIEW2_WEB_ERROR_STATUS_OPERATION_CANCELED: 14,
	COREWEBVIEW2_WEB_ERROR_STATUS_REDIRECT_FAILED: 15,
	COREWEBVIEW2_WEB_ERROR_STATUS_UNEXPECTED_ERROR: 16,
	COREWEBVIEW2_WEB_ERROR_STATUS_VALID_AUTHENTICATION_CREDENTIALS_REQUIRED: 17,
	COREWEBVIEW2_WEB_ERROR_STATUS_VALID_PROXY_AUTHENTICATION_REQUIRED: 18,
}

// enum COREWEBVIEW2_SCRIPT_DIALOG_KIND
var COREWEBVIEW2_SCRIPT_DIALOG_KIND = struct {
	COREWEBVIEW2_SCRIPT_DIALOG_KIND_ALERT int32
	COREWEBVIEW2_SCRIPT_DIALOG_KIND_CONFIRM int32
	COREWEBVIEW2_SCRIPT_DIALOG_KIND_PROMPT int32
	COREWEBVIEW2_SCRIPT_DIALOG_KIND_BEFOREUNLOAD int32
}{
	COREWEBVIEW2_SCRIPT_DIALOG_KIND_ALERT: 0,
	COREWEBVIEW2_SCRIPT_DIALOG_KIND_CONFIRM: 1,
	COREWEBVIEW2_SCRIPT_DIALOG_KIND_PROMPT: 2,
	COREWEBVIEW2_SCRIPT_DIALOG_KIND_BEFOREUNLOAD: 3,
}

// enum COREWEBVIEW2_PERMISSION_KIND
var COREWEBVIEW2_PERMISSION_KIND = struct {
	COREWEBVIEW2_PERMISSION_KIND_UNKNOWN_PERMISSION int32
	COREWEBVIEW2_PERMISSION_KIND_MICROPHONE int32
	COREWEBVIEW2_PERMISSION_KIND_CAMERA int32
	COREWEBVIEW2_PERMISSION_KIND_GEOLOCATION int32
	COREWEBVIEW2_PERMISSION_KIND_NOTIFICATIONS int32
	COREWEBVIEW2_PERMISSION_KIND_OTHER_SENSORS int32
	COREWEBVIEW2_PERMISSION_KIND_CLIPBOARD_READ int32
}{
	COREWEBVIEW2_PERMISSION_KIND_UNKNOWN_PERMISSION: 0,
	COREWEBVIEW2_PERMISSION_KIND_MICROPHONE: 1,
	COREWEBVIEW2_PERMISSION_KIND_CAMERA: 2,
	COREWEBVIEW2_PERMISSION_KIND_GEOLOCATION: 3,
	COREWEBVIEW2_PERMISSION_KIND_NOTIFICATIONS: 4,
	COREWEBVIEW2_PERMISSION_KIND_OTHER_SENSORS: 5,
	COREWEBVIEW2_PERMISSION_KIND_CLIPBOARD_READ: 6,
}

// enum COREWEBVIEW2_PERMISSION_STATE
var COREWEBVIEW2_PERMISSION_STATE = struct {
	COREWEBVIEW2_PERMISSION_STATE_DEFAULT int32
	COREWEBVIEW2_PERMISSION_STATE_ALLOW int32
	COREWEBVIEW2_PERMISSION_STATE_DENY int32
}{
	COREWEBVIEW2_PERMISSION_STATE_DEFAULT: 0,
	COREWEBVIEW2_PERMISSION_STATE_ALLOW: 1,
	COREWEBVIEW2_PERMISSION_STATE_DENY: 2,
}

// enum COREWEBVIEW2_PROCESS_FAILED_KIND
var COREWEBVIEW2_PROCESS_FAILED_KIND = struct {
	COREWEBVIEW2_PROCESS_FAILED_KIND_BROWSER_PROCESS_EXITED int32
	COREWEBVIEW2_PROCESS_FAILED_KIND_RENDER_PROCESS_EXITED int32
	COREWEBVIEW2_PROCESS_FAILED_KIND_RENDER_PROCESS_UNRESPONSIVE int32
	COREWEBVIEW2_PROCESS_FAILED_KIND_FRAME_RENDER_PROCESS_EXITED int32
	COREWEBVIEW2_PROCESS_FAILED_KIND_UTILITY_PROCESS_EXITED int32
	COREWEBVIEW2_PROCESS_FAILED_KIND_SANDBOX_HELPER_PROCESS_EXITED int32
	COREWEBVIEW2_PROCESS_FAILED_KIND_GPU_PROCESS_EXITED int32
	COREWEBVIEW2_PROCESS_FAILED_KIND_PPAPI_PLUGIN_PROCESS_EXITED int32
	COREWEBVIEW2_PROCESS_FAILED_KIND_PPAPI_BROKER_PROCESS_EXITED int32
	COREWEBVIEW2_PROCESS_FAILED_KIND_UNKNOWN_PROCESS_EXITED int32
}{
	COREWEBVIEW2_PROCESS_FAILED_KIND_BROWSER_PROCESS_EXITED: 0,
	COREWEBVIEW2_PROCESS_FAILED_KIND_RENDER_PROCESS_EXITED: 1,
	COREWEBVIEW2_PROCESS_FAILED_KIND_RENDER_PROCESS_UNRESPONSIVE: 2,
	COREWEBVIEW2_PROCESS_FAILED_KIND_FRAME_RENDER_PROCESS_EXITED: 3,
	COREWEBVIEW2_PROCESS_FAILED_KIND_UTILITY_PROCESS_EXITED: 4,
	COREWEBVIEW2_PROCESS_FAILED_KIND_SANDBOX_HELPER_PROCESS_EXITED: 5,
	COREWEBVIEW2_PROCESS_FAILED_KIND_GPU_PROCESS_EXITED: 6,
	COREWEBVIEW2_PROCESS_FAILED_KIND_PPAPI_PLUGIN_PROCESS_EXITED: 7,
	COREWEBVIEW2_PROCESS_FAILED_KIND_PPAPI_BROKER_PROCESS_EXITED: 8,
	COREWEBVIEW2_PROCESS_FAILED_KIND_UNKNOWN_PROCESS_EXITED: 9,
}

// enum COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT
var COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT = struct {
	COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT_PNG int32
	COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT_JPEG int32
}{
	COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT_PNG: 0,
	COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT_JPEG: 1,
}

// enum COREWEBVIEW2_WEB_RESOURCE_CONTEXT
var COREWEBVIEW2_WEB_RESOURCE_CONTEXT = struct {
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_ALL int32
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_DOCUMENT int32
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_STYLESHEET int32
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_IMAGE int32
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_MEDIA int32
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_FONT int32
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_SCRIPT int32
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_XML_HTTP_REQUEST int32
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_FETCH int32
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_TEXT_TRACK int32
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_EVENT_SOURCE int32
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_WEBSOCKET int32
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_MANIFEST int32
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_SIGNED_EXCHANGE int32
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_PING int32
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_CSP_VIOLATION_REPORT int32
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_OTHER int32
}{
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_ALL: 0,
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_DOCUMENT: 1,
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_STYLESHEET: 2,
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_IMAGE: 3,
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_MEDIA: 4,
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_FONT: 5,
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_SCRIPT: 6,
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_XML_HTTP_REQUEST: 7,
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_FETCH: 8,
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_TEXT_TRACK: 9,
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_EVENT_SOURCE: 10,
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_WEBSOCKET: 11,
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_MANIFEST: 12,
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_SIGNED_EXCHANGE: 13,
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_PING: 14,
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_CSP_VIOLATION_REPORT: 15,
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_OTHER: 16,
}

// enum COREWEBVIEW2_COOKIE_SAME_SITE_KIND
var COREWEBVIEW2_COOKIE_SAME_SITE_KIND = struct {
	COREWEBVIEW2_COOKIE_SAME_SITE_KIND_NONE int32
	COREWEBVIEW2_COOKIE_SAME_SITE_KIND_LAX int32
	COREWEBVIEW2_COOKIE_SAME_SITE_KIND_STRICT int32
}{
	COREWEBVIEW2_COOKIE_SAME_SITE_KIND_NONE: 0,
	COREWEBVIEW2_COOKIE_SAME_SITE_KIND_LAX: 1,
	COREWEBVIEW2_COOKIE_SAME_SITE_KIND_STRICT: 2,
}

// enum COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND
var COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND = struct {
	COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_DENY int32
	COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_ALLOW int32
	COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_DENY_CORS int32
}{
	COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_DENY: 0,
	COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_ALLOW: 1,
	COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_DENY_CORS: 2,
}

// enum COREWEBVIEW2_DOWNLOAD_STATE
var COREWEBVIEW2_DOWNLOAD_STATE = struct {
	COREWEBVIEW2_DOWNLOAD_STATE_IN_PROGRESS int32
	COREWEBVIEW2_DOWNLOAD_STATE_INTERRUPTED int32
	COREWEBVIEW2_DOWNLOAD_STATE_COMPLETED int32
}{
	COREWEBVIEW2_DOWNLOAD_STATE_IN_PROGRESS: 0,
	COREWEBVIEW2_DOWNLOAD_STATE_INTERRUPTED: 1,
	COREWEBVIEW2_DOWNLOAD_STATE_COMPLETED: 2,
}

// enum COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON
var COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = struct {
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NONE int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_FAILED int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_ACCESS_DENIED int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_NO_SPACE int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_NAME_TOO_LONG int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_TOO_LARGE int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_MALICIOUS int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_TRANSIENT_ERROR int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_BLOCKED_BY_POLICY int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_SECURITY_CHECK_FAILED int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_TOO_SHORT int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_HASH_MISMATCH int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_FAILED int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_TIMEOUT int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_DISCONNECTED int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_SERVER_DOWN int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_INVALID_REQUEST int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_FAILED int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_NO_RANGE int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_BAD_CONTENT int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_UNAUTHORIZED int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_CERTIFICATE_PROBLEM int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_FORBIDDEN int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_UNEXPECTED_RESPONSE int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_CONTENT_LENGTH_MISMATCH int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_CROSS_ORIGIN_REDIRECT int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_CANCELED int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_SHUTDOWN int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_PAUSED int32
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_DOWNLOAD_PROCESS_CRASHED int32
}{
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NONE: 0,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_FAILED: 1,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_ACCESS_DENIED: 2,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_NO_SPACE: 3,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_NAME_TOO_LONG: 4,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_TOO_LARGE: 5,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_MALICIOUS: 6,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_TRANSIENT_ERROR: 7,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_BLOCKED_BY_POLICY: 8,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_SECURITY_CHECK_FAILED: 9,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_TOO_SHORT: 10,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_HASH_MISMATCH: 11,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_FAILED: 12,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_TIMEOUT: 13,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_DISCONNECTED: 14,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_SERVER_DOWN: 15,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_INVALID_REQUEST: 16,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_FAILED: 17,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_NO_RANGE: 18,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_BAD_CONTENT: 19,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_UNAUTHORIZED: 20,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_CERTIFICATE_PROBLEM: 21,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_FORBIDDEN: 22,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_UNEXPECTED_RESPONSE: 23,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_CONTENT_LENGTH_MISMATCH: 24,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_CROSS_ORIGIN_REDIRECT: 25,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_CANCELED: 26,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_SHUTDOWN: 27,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_PAUSED: 28,
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_DOWNLOAD_PROCESS_CRASHED: 29,
}

// enum COREWEBVIEW2_CLIENT_CERTIFICATE_KIND
var COREWEBVIEW2_CLIENT_CERTIFICATE_KIND = struct {
	COREWEBVIEW2_CLIENT_CERTIFICATE_KIND_SMART_CARD int32
	COREWEBVIEW2_CLIENT_CERTIFICATE_KIND_PIN int32
	COREWEBVIEW2_CLIENT_CERTIFICATE_KIND_OTHER int32
}{
	COREWEBVIEW2_CLIENT_CERTIFICATE_KIND_SMART_CARD: 0,
	COREWEBVIEW2_CLIENT_CERTIFICATE_KIND_PIN: 1,
	COREWEBVIEW2_CLIENT_CERTIFICATE_KIND_OTHER: 2,
}

// enum COREWEBVIEW2_PRINT_ORIENTATION
var COREWEBVIEW2_PRINT_ORIENTATION = struct {
	COREWEBVIEW2_PRINT_ORIENTATION_PORTRAIT int32
	COREWEBVIEW2_PRINT_ORIENTATION_LANDSCAPE int32
}{
	COREWEBVIEW2_PRINT_ORIENTATION_PORTRAIT: 0,
	COREWEBVIEW2_PRINT_ORIENTATION_LANDSCAPE: 1,
}

// enum COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT
var COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT = struct {
	COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT_TOP_LEFT int32
	COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT_TOP_RIGHT int32
	COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT_BOTTOM_LEFT int32
	COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT_BOTTOM_RIGHT int32
}{
	COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT_TOP_LEFT: 0,
	COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT_TOP_RIGHT: 1,
	COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT_BOTTOM_LEFT: 2,
	COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT_BOTTOM_RIGHT: 3,
}

// enum COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND
var COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND = struct {
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_COMMAND int32
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_CHECK_BOX int32
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_RADIO int32
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_SEPARATOR int32
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_SUBMENU int32
}{
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_COMMAND: 0,
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_CHECK_BOX: 1,
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_RADIO: 2,
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_SEPARATOR: 3,
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_SUBMENU: 4,
}

// enum COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND
var COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND = struct {
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_PAGE int32
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_IMAGE int32
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_SELECTED_TEXT int32
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_AUDIO int32
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_VIDEO int32
}{
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_PAGE: 0,
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_IMAGE: 1,
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_SELECTED_TEXT: 2,
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_AUDIO: 3,
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_VIDEO: 4,
}

// enum COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND
var COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND = struct {
	COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND_NORMAL int32
	COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND_FAILED int32
}{
	COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND_NORMAL: 0,
	COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND_FAILED: 1,
}

// enum COREWEBVIEW2_MOUSE_EVENT_KIND
var COREWEBVIEW2_MOUSE_EVENT_KIND = struct {
	COREWEBVIEW2_MOUSE_EVENT_KIND_HORIZONTAL_WHEEL int32
	COREWEBVIEW2_MOUSE_EVENT_KIND_LEFT_BUTTON_DOUBLE_CLICK int32
	COREWEBVIEW2_MOUSE_EVENT_KIND_LEFT_BUTTON_DOWN int32
	COREWEBVIEW2_MOUSE_EVENT_KIND_LEFT_BUTTON_UP int32
	COREWEBVIEW2_MOUSE_EVENT_KIND_LEAVE int32
	COREWEBVIEW2_MOUSE_EVENT_KIND_MIDDLE_BUTTON_DOUBLE_CLICK int32
	COREWEBVIEW2_MOUSE_EVENT_KIND_MIDDLE_BUTTON_DOWN int32
	COREWEBVIEW2_MOUSE_EVENT_KIND_MIDDLE_BUTTON_UP int32
	COREWEBVIEW2_MOUSE_EVENT_KIND_MOVE int32
	COREWEBVIEW2_MOUSE_EVENT_KIND_RIGHT_BUTTON_DOUBLE_CLICK int32
	COREWEBVIEW2_MOUSE_EVENT_KIND_RIGHT_BUTTON_DOWN int32
	COREWEBVIEW2_MOUSE_EVENT_KIND_RIGHT_BUTTON_UP int32
	COREWEBVIEW2_MOUSE_EVENT_KIND_WHEEL int32
	COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_DOUBLE_CLICK int32
	COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_DOWN int32
	COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_UP int32
}{
	COREWEBVIEW2_MOUSE_EVENT_KIND_HORIZONTAL_WHEEL: 526,
	COREWEBVIEW2_MOUSE_EVENT_KIND_LEFT_BUTTON_DOUBLE_CLICK: 515,
	COREWEBVIEW2_MOUSE_EVENT_KIND_LEFT_BUTTON_DOWN: 513,
	COREWEBVIEW2_MOUSE_EVENT_KIND_LEFT_BUTTON_UP: 514,
	COREWEBVIEW2_MOUSE_EVENT_KIND_LEAVE: 675,
	COREWEBVIEW2_MOUSE_EVENT_KIND_MIDDLE_BUTTON_DOUBLE_CLICK: 521,
	COREWEBVIEW2_MOUSE_EVENT_KIND_MIDDLE_BUTTON_DOWN: 519,
	COREWEBVIEW2_MOUSE_EVENT_KIND_MIDDLE_BUTTON_UP: 520,
	COREWEBVIEW2_MOUSE_EVENT_KIND_MOVE: 512,
	COREWEBVIEW2_MOUSE_EVENT_KIND_RIGHT_BUTTON_DOUBLE_CLICK: 518,
	COREWEBVIEW2_MOUSE_EVENT_KIND_RIGHT_BUTTON_DOWN: 516,
	COREWEBVIEW2_MOUSE_EVENT_KIND_RIGHT_BUTTON_UP: 517,
	COREWEBVIEW2_MOUSE_EVENT_KIND_WHEEL: 522,
	COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_DOUBLE_CLICK: 525,
	COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_DOWN: 523,
	COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_UP: 524,
}

// enum COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS
var COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS = struct {
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_NONE int32
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_LEFT_BUTTON int32
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_RIGHT_BUTTON int32
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_SHIFT int32
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_CONTROL int32
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_MIDDLE_BUTTON int32
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_X_BUTTON1 int32
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_X_BUTTON2 int32
}{
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_NONE: 0,
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_LEFT_BUTTON: 1,
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_RIGHT_BUTTON: 2,
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_SHIFT: 4,
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_CONTROL: 8,
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_MIDDLE_BUTTON: 16,
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_X_BUTTON1: 32,
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_X_BUTTON2: 64,
}

// enum COREWEBVIEW2_POINTER_EVENT_KIND
var COREWEBVIEW2_POINTER_EVENT_KIND = struct {
	COREWEBVIEW2_POINTER_EVENT_KIND_ACTIVATE int32
	COREWEBVIEW2_POINTER_EVENT_KIND_DOWN int32
	COREWEBVIEW2_POINTER_EVENT_KIND_ENTER int32
	COREWEBVIEW2_POINTER_EVENT_KIND_LEAVE int32
	COREWEBVIEW2_POINTER_EVENT_KIND_UP int32
	COREWEBVIEW2_POINTER_EVENT_KIND_UPDATE int32
}{
	COREWEBVIEW2_POINTER_EVENT_KIND_ACTIVATE: 587,
	COREWEBVIEW2_POINTER_EVENT_KIND_DOWN: 582,
	COREWEBVIEW2_POINTER_EVENT_KIND_ENTER: 585,
	COREWEBVIEW2_POINTER_EVENT_KIND_LEAVE: 586,
	COREWEBVIEW2_POINTER_EVENT_KIND_UP: 583,
	COREWEBVIEW2_POINTER_EVENT_KIND_UPDATE: 581,
}

// enum COREWEBVIEW2_BOUNDS_MODE
var COREWEBVIEW2_BOUNDS_MODE = struct {
	COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS int32
	COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE int32
}{
	COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS: 0,
	COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE: 1,
}

// enum COREWEBVIEW2_PROCESS_KIND
var COREWEBVIEW2_PROCESS_KIND = struct {
	COREWEBVIEW2_PROCESS_KIND_BROWSER int32
	COREWEBVIEW2_PROCESS_KIND_RENDERER int32
	COREWEBVIEW2_PROCESS_KIND_UTILITY int32
	COREWEBVIEW2_PROCESS_KIND_SANDBOX_HELPER int32
	COREWEBVIEW2_PROCESS_KIND_GPU int32
	COREWEBVIEW2_PROCESS_KIND_PPAPI_PLUGIN int32
	COREWEBVIEW2_PROCESS_KIND_PPAPI_BROKER int32
}{
	COREWEBVIEW2_PROCESS_KIND_BROWSER: 0,
	COREWEBVIEW2_PROCESS_KIND_RENDERER: 1,
	COREWEBVIEW2_PROCESS_KIND_UTILITY: 2,
	COREWEBVIEW2_PROCESS_KIND_SANDBOX_HELPER: 3,
	COREWEBVIEW2_PROCESS_KIND_GPU: 4,
	COREWEBVIEW2_PROCESS_KIND_PPAPI_PLUGIN: 5,
	COREWEBVIEW2_PROCESS_KIND_PPAPI_BROKER: 6,
}

// enum COREWEBVIEW2_PROCESS_FAILED_REASON
var COREWEBVIEW2_PROCESS_FAILED_REASON = struct {
	COREWEBVIEW2_PROCESS_FAILED_REASON_UNEXPECTED int32
	COREWEBVIEW2_PROCESS_FAILED_REASON_UNRESPONSIVE int32
	COREWEBVIEW2_PROCESS_FAILED_REASON_TERMINATED int32
	COREWEBVIEW2_PROCESS_FAILED_REASON_CRASHED int32
	COREWEBVIEW2_PROCESS_FAILED_REASON_LAUNCH_FAILED int32
	COREWEBVIEW2_PROCESS_FAILED_REASON_OUT_OF_MEMORY int32
}{
	COREWEBVIEW2_PROCESS_FAILED_REASON_UNEXPECTED: 0,
	COREWEBVIEW2_PROCESS_FAILED_REASON_UNRESPONSIVE: 1,
	COREWEBVIEW2_PROCESS_FAILED_REASON_TERMINATED: 2,
	COREWEBVIEW2_PROCESS_FAILED_REASON_CRASHED: 3,
	COREWEBVIEW2_PROCESS_FAILED_REASON_LAUNCH_FAILED: 4,
	COREWEBVIEW2_PROCESS_FAILED_REASON_OUT_OF_MEMORY: 5,
}

// enum COREWEBVIEW2_PREFERRED_COLOR_SCHEME
var COREWEBVIEW2_PREFERRED_COLOR_SCHEME = struct {
	COREWEBVIEW2_PREFERRED_COLOR_SCHEME_AUTO int32
	COREWEBVIEW2_PREFERRED_COLOR_SCHEME_LIGHT int32
	COREWEBVIEW2_PREFERRED_COLOR_SCHEME_DARK int32
}{
	COREWEBVIEW2_PREFERRED_COLOR_SCHEME_AUTO: 0,
	COREWEBVIEW2_PREFERRED_COLOR_SCHEME_LIGHT: 1,
	COREWEBVIEW2_PREFERRED_COLOR_SCHEME_DARK: 2,
}

// enum COREWEBVIEW2_PDF_TOOLBAR_ITEMS
var COREWEBVIEW2_PDF_TOOLBAR_ITEMS = struct {
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_NONE int32
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_SAVE int32
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_PRINT int32
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_SAVE_AS int32
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_ZOOM_IN int32
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_ZOOM_OUT int32
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_ROTATE int32
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_FIT_PAGE int32
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_PAGE_LAYOUT int32
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_BOOKMARKS int32
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_PAGE_SELECTOR int32
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_SEARCH int32
}{
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_NONE: 0,
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_SAVE: 1,
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_PRINT: 2,
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_SAVE_AS: 4,
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_ZOOM_IN: 8,
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_ZOOM_OUT: 16,
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_ROTATE: 32,
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_FIT_PAGE: 64,
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_PAGE_LAYOUT: 128,
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_BOOKMARKS: 256,
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_PAGE_SELECTOR: 512,
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_SEARCH: 1024,
}
