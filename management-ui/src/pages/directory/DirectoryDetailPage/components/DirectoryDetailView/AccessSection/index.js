// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { number, shape, func } from 'prop-types'
import { observer } from 'mobx-react'
import { Spinner, Button } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import { outgoingAccessRequestPropTypes } from '../../../../../../models/OutgoingAccessRequestModel'
import { accessProofPropTypes } from '../../../../../../models/AccessProofModel'
import {
  SHOW_REQUEST_ACCESS,
  SHOW_HAS_ACCESS,
  SHOW_REQUEST_CREATED,
  SHOW_REQUEST_FAILED,
  SHOW_REQUEST_RECEIVED,
  SHOW_REQUEST_REJECTED,
  SHOW_ACCESS_REVOKED,
} from '../../../../directoryServiceAccessState'
import {
  IconWarningCircleFill,
  IconKey,
  IconCheck,
} from '../../../../../../icons'
import Switch from '../../../../../../components/Switch'
import { Section, IconItem, StateDetail, ErrorText } from './index.styles'

const AccessSection = ({
  displayState,
  latestAccessRequest,
  latestAccessProof,
  requestAccess,
}) => {
  const { t } = useTranslation()

  const onRequestAccessButtonClick = (evt) => {
    evt.stopPropagation()
    requestAccess()
  }

  return (
    <Section data-testid="request-access-section">
      <Switch test={displayState}>
        <Switch.Case value={SHOW_REQUEST_ACCESS}>
          {() => (
            <>
              <IconItem as={IconKey} />
              <StateDetail>{t('You have no access')}</StateDetail>
              <Button onClick={onRequestAccessButtonClick}>
                {t('Request Access')}
              </Button>
            </>
          )}
        </Switch.Case>

        <Switch.Case value={SHOW_REQUEST_FAILED}>
          {() => (
            <>
              <IconItem as={IconKey} />
              <StateDetail>
                <span>{t('Access request')}</span>
                <ErrorText>
                  <IconWarningCircleFill title={t('Error')} />
                  {t('Request could not be sent')}
                </ErrorText>
              </StateDetail>
            </>
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
              <IconItem as={IconKey} />
              <StateDetail>
                <span>{t('Access requested')}</span>
                <small>
                  {t('On date', { date: latestAccessRequest.updatedAt })}
                </small>
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
                  {t('Since date', {
                    date: latestAccessProof.createdAt,
                  })}
                </small>
              </StateDetail>
            </>
          )}
        </Switch.Case>

        <Switch.Case value={SHOW_REQUEST_REJECTED}>
          {() => (
            <>
              <IconItem as={IconKey} />
              <StateDetail>
                <span>{t('Access request rejected')}</span>
                <small>
                  {t('On date', {
                    date: latestAccessRequest.updatedAt,
                  })}
                </small>
              </StateDetail>
            </>
          )}
        </Switch.Case>

        <Switch.Case value={SHOW_ACCESS_REVOKED}>
          {() => (
            <>
              <IconItem as={IconKey} />
              <StateDetail>
                <span>{t('Your access was revoked')}</span>
                <small>
                  {t(
                    'Due to a technical limitation you are not yet able to request access again',
                  )}
                </small>
              </StateDetail>
            </>
          )}
        </Switch.Case>
      </Switch>
    </Section>
  )
}

AccessSection.propTypes = {
  displayState: number,
  latestAccessRequest: shape(outgoingAccessRequestPropTypes),
  latestAccessProof: shape(accessProofPropTypes),
  requestAccess: func,
}

export default observer(AccessSection)
