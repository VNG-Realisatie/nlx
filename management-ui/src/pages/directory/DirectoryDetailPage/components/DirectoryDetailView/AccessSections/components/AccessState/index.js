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
  SHOW_REQUEST_FAILED,
  SHOW_REQUEST_RECEIVED,
  SHOW_REQUEST_REJECTED,
  SHOW_ACCESS_REVOKED,
  SHOW_REQUEST_WITHDRAWN,
  SHOW_ACCESS_TERMINATED,
} from '../../../../../../directoryServiceAccessState'
import { IconCheck, IconClose } from '../../../../../../../../icons'
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
  onWithdrawAccessButtonClick,
}) => {
  const { t } = useTranslation()

  let displayState = getDirectoryServiceAccessUIState(
    accessRequest,
    accessProof,
  )

  const onRequestAccessButtonClick = (event) => {
    event.stopPropagation()
    onRequestAccess()
  }

  return (
    <section data-testid="request-access-section">
      {isLoading ? (
        <StateContainer>
          <IconItem as={Spinner} />
          <StateDetail>
            <span>{t('Processing request')}…</span>
          </StateDetail>
        </StateContainer>
      ) : (
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
                {t(accessRequest.errorDetails.cause)}
              </StyledAlert>
            )}
          </Switch.Case>

          <Switch.Case value={SHOW_REQUEST_WITHDRAWN}>
            {() => (
              <StateContainer>
                <StateDetail>
                  <span>{t('Latest access request has been withdrawn')}</span>
                  <small>
                    {t('On date', { date: accessRequest.updatedAt })}
                  </small>
                </StateDetail>

                <Button onClick={onRequestAccessButtonClick}>
                  {t('Request access')}
                </Button>
              </StateContainer>
            )}
          </Switch.Case>

          <Switch.Case value={SHOW_ACCESS_TERMINATED}>
            {() => (
              <StateContainer>
                <StateDetail>
                  <span>{t('Latest access request has been terminated')}</span>
                  <small>
                    {t('On date', { date: accessProof.terminatedAt })}
                  </small>
                </StateDetail>

                <Button onClick={onRequestAccessButtonClick}>
                  {t('Request access')}
                </Button>
              </StateContainer>
            )}
          </Switch.Case>

          <Switch.Case value={SHOW_REQUEST_RECEIVED}>
            {() => (
              <StateContainer>
                <StateDetail>
                  <span>{t('Access requested')}</span>
                  <small>
                    {t('On date', { date: accessRequest.updatedAt })}
                  </small>
                </StateDetail>

                <Button
                  onClick={onWithdrawAccessButtonClick}
                  variant="link"
                  title={t(
                    'Withdraw access request for Outways with public key fingerprint {{publicKeyFingerprint}}',
                    {
                      publicKeyFingerprint: accessRequest.publicKeyFingerprint,
                    },
                  )}
                >
                  <IconClose inline />
                  {t('Withdraw')}
                </Button>
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
      )}
    </section>
  )
}

AccessState.propTypes = {
  isLoading: bool,
  accessRequest: instanceOf(OutgoingAccessRequestModel),
  accessProof: instanceOf(AccessProofModel),
  onRequestAccess: func,
  onRetryRequestAccess: func,
  onWithdrawAccessButtonClick: func,
}

export default observer(AccessState)
