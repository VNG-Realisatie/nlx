// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { func, instanceOf } from 'prop-types'
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
} from '../../../../directoryServiceAccessState'
import { IconCheck } from '../../../../../../icons'
import Switch from '../../../../../../components/Switch'
import OutgoingAccessRequestModel from '../../../../../../stores/models/OutgoingAccessRequestModel'
import AccessProofModel from '../../../../../../stores/models/AccessProofModel'
import { StyledAlert } from '../index.styles'
import { Section, IconItem, StateDetail } from './index.styles'

const AccessSection = ({
  accessRequest,
  accessProof,
  onRequestAccess,
  onRetryRequestAccess,
}) => {
  const { t } = useTranslation()

  const displayState = getDirectoryServiceAccessUIState(
    accessRequest,
    accessProof,
  )

  const onRequestAccessButtonClick = (event) => {
    event.stopPropagation()
    onRequestAccess()
  }

  return (
    <Section data-testid="request-access-section">
      <Switch test={displayState}>
        <Switch.Case value={SHOW_REQUEST_ACCESS}>
          {() => (
            <>
              <StateDetail>{t('You have no access')}</StateDetail>
              <Button onClick={onRequestAccessButtonClick}>
                {t('Request access')}
              </Button>
            </>
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
            <>
              <IconItem as={Spinner} />
              <StateDetail>
                <span>{t('Sending request')}…</span>
              </StateDetail>
            </>
          )}
        </Switch.Case>

        <Switch.Case value={SHOW_REQUEST_RECEIVED}>
          {() => (
            <>
              <StateDetail>
                <span>{t('Access requested')}</span>
                <small>{t('On date', { date: accessRequest.updatedAt })}</small>
              </StateDetail>
            </>
          )}
        </Switch.Case>

        <Switch.Case value={SHOW_HAS_ACCESS}>
          {() => (
            <>
              <IconItem as={IconCheck} />
              <StateDetail>
                <span>{t('You have access')}</span>
                <small>
                  {t('Since datetime', {
                    date: accessProof.createdAt,
                  })}
                </small>
              </StateDetail>
            </>
          )}
        </Switch.Case>

        <Switch.Case value={SHOW_REQUEST_REJECTED}>
          {() => (
            <>
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
            </>
          )}
        </Switch.Case>

        <Switch.Case value={SHOW_ACCESS_REVOKED}>
          {() => (
            <>
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
            </>
          )}
        </Switch.Case>
        <Switch.Default>{() => null}</Switch.Default>
      </Switch>
    </Section>
  )
}

AccessSection.propTypes = {
  accessRequest: instanceOf(OutgoingAccessRequestModel),
  accessProof: instanceOf(AccessProofModel),
  onRequestAccess: func,
  onRetryRequestAccess: func,
}

export default observer(AccessSection)
