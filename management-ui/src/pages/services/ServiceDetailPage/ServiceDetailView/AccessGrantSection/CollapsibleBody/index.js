// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { useContext } from 'react'
import { Table, ToasterContext } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { array, bool, func } from 'prop-types'
import { StyledCollapsibleBody } from '../../../../../../components/DetailView'
import UpdateUiButton from '../../UpdateUiButton'
import AccessGrantRow from './AccessGrantRow'

const CollapsibleBody = ({
  accessGrants,
  showLoadIncomingDataButton,
  onClickLoadIncomingDataHandler,
}) => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)

  const revokeHandler = async (accessGrant) => {
    await accessGrant.revoke()

    if (!accessGrant.error) {
      showToast({
        title: t('Access revoked'),
        variant: 'success',
      })
    } else {
      showToast({
        title: t('Failed to revoke access grant'),
        body: t('Please try again'),
        variant: 'error',
      })
      console.error(accessGrant.error)
    }
  }

  return (
    <StyledCollapsibleBody>
      {accessGrants.length ? (
        <Table data-testid="service-accessgrant-list">
          <tbody>
            {accessGrants.map((accessGrant) => (
              <AccessGrantRow
                key={accessGrant.id}
                accessGrant={accessGrant}
                revokeHandler={revokeHandler}
              />
            ))}
          </tbody>
        </Table>
      ) : (
        <small>{t('There are no organizations with access')}</small>
      )}

      {showLoadIncomingDataButton ? (
        <UpdateUiButton onClick={onClickLoadIncomingDataHandler}>
          Nieuwe verzoeken
        </UpdateUiButton>
      ) : null}
    </StyledCollapsibleBody>
  )
}

CollapsibleBody.propTypes = {
  accessGrants: array,
  showLoadIncomingDataButton: bool,
  onClickLoadIncomingDataHandler: func,
}

const noop = () => {}

CollapsibleBody.defaultProps = {
  accessGrants: [],
  showLoadIncomingDataButton: false,
  onClickLoadIncomingDataHandler: noop,
}

export default CollapsibleBody
