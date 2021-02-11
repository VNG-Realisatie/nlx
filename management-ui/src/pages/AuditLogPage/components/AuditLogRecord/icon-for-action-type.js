// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import {
  ACTION_ACCESS_GRANT_REVOKE,
  ACTION_INCOMING_ACCESS_REQUEST_ACCEPT,
  ACTION_INCOMING_ACCESS_REQUEST_REJECT,
  ACTION_INSIGHT_CONFIGURATION_UPDATE,
  ACTION_LOGIN_FAIL,
  ACTION_LOGIN_SUCCESS,
  ACTION_LOGOUT_SUCCESS,
  ACTION_ORGANIZATION_SETTINGS_UPDATE,
  ACTION_OUTGOING_ACCESS_REQUEST_CREATE,
  ACTION_SERVICE_CREATE,
  ACTION_SERVICE_DELETE,
  ACTION_SERVICE_UPDATE,
} from '../../../../stores/models/AuditLogModel'
import {
  IconCheck,
  IconClose,
  IconKey,
  IconRevoke,
  IconServices,
  IconSettings,
  IconShutdown,
  IconWarningCircle,
} from '../../../../icons'

const iconForActionType = (actionType) => {
  switch (actionType) {
    case ACTION_LOGIN_SUCCESS:
    case ACTION_LOGOUT_SUCCESS:
    case ACTION_LOGIN_FAIL:
      return IconShutdown

    case ACTION_INCOMING_ACCESS_REQUEST_ACCEPT:
      return IconCheck

    case ACTION_INCOMING_ACCESS_REQUEST_REJECT:
      return IconClose

    case ACTION_ACCESS_GRANT_REVOKE:
      return IconRevoke

    case ACTION_OUTGOING_ACCESS_REQUEST_CREATE:
      return IconKey

    case ACTION_SERVICE_CREATE:
    case ACTION_SERVICE_UPDATE:
    case ACTION_SERVICE_DELETE:
      return IconServices

    case ACTION_ORGANIZATION_SETTINGS_UPDATE:
    case ACTION_INSIGHT_CONFIGURATION_UPDATE:
      return IconSettings

    default:
      return IconWarningCircle
  }
}

export default iconForActionType
