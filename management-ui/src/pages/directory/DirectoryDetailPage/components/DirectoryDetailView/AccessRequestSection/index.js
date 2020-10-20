// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, func } from 'prop-types'
import { observer } from 'mobx-react'
import { Spinner, Button } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import {
  outgoingAccessRequestPropTypes,
  ACCESS_REQUEST_STATES,
} from '../../../../../../models/OutgoingAccessRequestModel'
import {
  IconWarningCircleFill,
  IconKey,
  IconCheck,
} from '../../../../../../icons'
import { AccessSection, IconItem, StateDetail, ErrorText } from './index.styles'

const { FAILED, CREATED, RECEIVED, APPROVED, REJECTED } = ACCESS_REQUEST_STATES

const getStateUI = (latestAccessRequest, t) => {
  switch (latestAccessRequest.state) {
    case FAILED:
      return (
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
      )

    case CREATED:
      return (
        <>
          <IconItem as={Spinner} />
          <StateDetail>
            <span>{t('Sending request')}…</span>
          </StateDetail>
        </>
      )

    case RECEIVED:
      return (
        <>
          <IconItem as={IconKey} />
          <StateDetail>
            <span>{t('Access requested')}</span>
            <small>
              {t('On date', { date: new Date(latestAccessRequest.updatedAt) })}
            </small>
          </StateDetail>
        </>
      )

    case APPROVED:
      return (
        <>
          <IconItem as={IconCheck} />
          <StateDetail>
            <span>{t('You have access')}</span>
            <small>
              {t('Since date', {
                date: new Date(latestAccessRequest.updatedAt),
              })}
            </small>
          </StateDetail>
        </>
      )

    case REJECTED:
      return (
        <>
          <IconItem as={IconKey} />
          <StateDetail>
            <span>{t('Access request rejected')}</span>
            <small>
              {t('On date', {
                date: new Date(latestAccessRequest.updatedAt),
              })}
            </small>
          </StateDetail>
        </>
      )

    default:
      throw new Error(
        `Can not determine message for unknown state '${latestAccessRequest.state}'`,
      )
  }
}

const AccessRequestSection = ({ latestAccessRequest, requestAccess }) => {
  const { t } = useTranslation()

  const onRequestAccessButtonClick = (evt) => {
    evt.stopPropagation()
    requestAccess()
  }

  return (
    <AccessSection data-testid="request-access-section">
      {latestAccessRequest ? (
        getStateUI(latestAccessRequest, t)
      ) : (
        <>
          <IconItem as={IconKey} />
          <StateDetail>{t('You have no access')}</StateDetail>
          <Button onClick={onRequestAccessButtonClick}>
            {t('Request Access')}
          </Button>
        </>
      )}
    </AccessSection>
  )
}

AccessRequestSection.propTypes = {
  latestAccessRequest: shape(outgoingAccessRequestPropTypes),
  requestAccess: func,
}

export default observer(AccessRequestSection)
