// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import React from 'react'
import { useTranslation } from 'react-i18next'
import { Collapsible } from '@commonground/design-system'
import { arrayOf, instanceOf } from 'prop-types'
import { IconServices } from '../../../../../../../icons'
import Amount from '../../../../../../../components/Amount'
import Table from '../../../../../../../components/Table'
import {
  DetailHeading,
  StyledCollapsibleBody,
} from '../../../../../../../components/DetailView'
import AccessProofModel from '../../../../../../../stores/models/AccessProofModel'
import { OrganizationName, Separator } from './index.styles'

const AccessProofs = ({ accessProofs }) => {
  const { t } = useTranslation()
  return (
    <Collapsible
      title={
        <DetailHeading>
          <IconServices />
          {t('Requestable services')}
          <Amount value={accessProofs.length} />
        </DetailHeading>
      }
      ariaLabel={t('Requestable services')}
      buttonLabels={{
        open: t('Open'),
        close: t('Close'),
      }}
    >
      <StyledCollapsibleBody>
        {accessProofs.length ? (
          <Table role="grid" withLinks>
            <tbody>
              {accessProofs.map(
                ({ serviceName, organization, publicKeyFingerprint }) => (
                  <Table.Tr
                    name=""
                    to={`/directory/${organization.serialNumber}/${serviceName}`}
                    key={`${organization.serialNumber}_${serviceName}`}
                  >
                    <Table.Td>
                      <OrganizationName>{organization.name}</OrganizationName>
                      <Separator> - </Separator>
                      <small>{serviceName}</small>
                      <br />
                      <small>{publicKeyFingerprint}</small>
                    </Table.Td>
                  </Table.Tr>
                ),
              )}
            </tbody>
          </Table>
        ) : (
          <small>{t('No services have been connected')}</small>
        )}
      </StyledCollapsibleBody>
    </Collapsible>
  )
}

AccessProofs.propTypes = {
  accessProofs: arrayOf(instanceOf(AccessProofModel)),
}

AccessProofs.defaultProps = {
  accessProofs: [],
}

export default AccessProofs
