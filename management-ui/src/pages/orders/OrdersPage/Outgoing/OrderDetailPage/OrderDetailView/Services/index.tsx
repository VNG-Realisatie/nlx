// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import React from 'react'
import { useTranslation } from 'react-i18next'
import { Collapsible } from '@commonground/design-system'
import { IconServices } from '../../../../../../../icons'
import Amount from '../../../../../../../components/Amount'
import Table from '../../../../../../../components/Table'
import {
  DetailHeading,
  StyledCollapsibleBody,
} from '../../../../../../../components/DetailView'
import Service from '../../../../../../../types/Service'
import { OrganizationName, Separator } from './index.styles'

const Services: React.FC<{ services: Service[] }> = ({ services }) => {
  const { t } = useTranslation()

  const ConnectedServices = () => {
    if (services.length) {
      return (
        <Table role="grid" withLinks>
          <tbody>
            {services.map(({ service, organization }) => (
              <Table.Tr
                name=""
                to={`/directory/${organization.serialNumber}/${service}`}
                key={`${organization.serialNumber}_${service}`}
              >
                <Table.Td>
                  <OrganizationName>
                    {organization.name} ({organization.serialNumber})
                  </OrganizationName>
                  <Separator> - </Separator>
                  <small>{service}</small>
                </Table.Td>
              </Table.Tr>
            ))}
          </tbody>
        </Table>
      )
    }
    return <small>{t('No services have been connected')}</small>
  }

  return (
    <Collapsible
      title={
        <DetailHeading>
          <IconServices />
          {t('Requestable services')}
          <Amount value={services.length} />
        </DetailHeading>
      }
      ariaLabel={t('Requestable services')}
      buttonLabels={{
        open: t('Open'),
        close: t('Close'),
      }}
    >
      <StyledCollapsibleBody>
        <ConnectedServices />
      </StyledCollapsibleBody>
    </Collapsible>
  )
}

export default Services
