// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { bool, func, instanceOf } from 'prop-types'
import { observer } from 'mobx-react'
import { Spinner, Button, Alert } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import getDirectoryServiceAccessUIState, {
  SHOW_REQUEST_ACCESS,
  SHOW_HAS_ACCESS,
  SHOW_REQUEST_CREATED,
  SHOW_REQUEST_FAILED,
  SHOW_REQUEST_RECEIVED,
  SHOW_REQUEST_REJECTED,
  SHOW_ACCESS_REVOKED,
} from '../../../../../../directoryServiceAccessState'
import { IconCheck } from '../../../../../../../../icons'
import Switch from '../../../../../../../../components/Switch'
import OutgoingAccessRequestModel from '../../../../../../../../stores/models/OutgoingAccessRequestModel'
import AccessProofModel from '../../../../../../../../stores/models/AccessProofModel'
import {
  IconItem,
  StateDetail,
  StyledAlert,
  StateContainer,
} from './index.styles'

const AccessState = ({
  isLoading,
  accessRequest,
  accessProof,
  onRequestAccess,
  onRetryRequestAccess,
}) => {
  const { t } = useTranslation()

  let displayState = getDirectoryServiceAccessUIState(
    accessRequest,
    accessProof,
  )

  if (isLoading) {
    displayState = SHOW_REQUEST_CREATED
  }

  const onRequestAccessButtonClick = (event) => {
    event.stopPropagation()
    onRequestAccess()
  }

  return (
    <section data-testid="request-access-section">
      <Switch test={displayState}>
        <Switch.Case value={SHOW_REQUEST_ACCESS}>
          {() => (
            <StateContainer>
              <StateDetail>{t('You have no access')}</StateDetail>
              <Button onClick={onRequestAccessButtonClick}>
                {t('Request access')}
              </Button>
            </StateContainer>
          )}
        </Switch.Case>

        <Switch.Case value={SHOW_REQUEST_FAILED}>
          {() => (
            <StyledAlert
              variant="error"
              title={t('Request could not be sent')}
              actions={[
                <Alert.ActionButton
                  key="send-request-access-action-button"
                  onClick={onRetryRequestAccess}
                >
                  {t('Retry')}
                </Alert.ActionButton>,
              ]}
            >
              {accessRequest.errorDetails.cause}
            </StyledAlert>
          )}
        </Switch.Case>

        <Switch.Case value={SHOW_REQUEST_CREATED}>
          {() => (
            <StateContainer>
              <IconItem as={Spinner} />
              <StateDetail>
                <span>{t('Sending request')}…</span>
              </StateDetail>
            </StateContainer>
          )}
        </Switch.Case>

        <Switch.Case value={SHOW_REQUEST_RECEIVED}>
          {() => (
            <StateContainer>
              <StateDetail>
                <span>{t('Access requested')}</span>
                <small>{t('On date', { date: accessRequest.updatedAt })}</small>
              </StateDetail>
            </StateContainer>
          )}
        </Switch.Case>

        <Switch.Case value={SHOW_HAS_ACCESS}>
          {() => (
            <StateContainer>
              <IconItem as={IconCheck} />
              <StateDetail>
                <span>{t('You have access')}</span>
                <small>
                  {t('Since datetime', {
                    date: accessProof.createdAt,
                  })}
                </small>
              </StateDetail>
            </StateContainer>
          )}
        </Switch.Case>

        <Switch.Case value={SHOW_REQUEST_REJECTED}>
          {() => (
            <StateContainer>
              <StateDetail>
                <span>{t('Access request rejected')}</span>
                <small>
                  {t('On date', {
                    date: accessRequest.updatedAt,
                  })}
                </small>
              </StateDetail>
              <Button onClick={onRequestAccessButtonClick}>
                {t('Request access')}
              </Button>
            </StateContainer>
          )}
        </Switch.Case>

        <Switch.Case value={SHOW_ACCESS_REVOKED}>
          {() => (
            <StateContainer>
              <StateDetail>
                <span>{t('Your access was revoked')}</span>
                <small>
                  {t('On date', {
                    date: accessProof.revokedAt,
                  })}
                </small>
              </StateDetail>
              <Button onClick={onRequestAccessButtonClick}>
                {t('Request access')}
              </Button>
            </StateContainer>
          )}
        </Switch.Case>
        <Switch.Default>{() => null}</Switch.Default>
      </Switch>
    </section>
  )
}

AccessState.propTypes = {
  isLoading: bool,
  accessRequest: instanceOf(OutgoingAccessRequestModel),
  accessProof: instanceOf(AccessProofModel),
  onRequestAccess: func,
  onRetryRequestAccess: func,
}

export default observer(AccessState)
