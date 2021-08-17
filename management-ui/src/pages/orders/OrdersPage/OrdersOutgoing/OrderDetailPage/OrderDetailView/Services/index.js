// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import React from 'react'
import { arrayOf, object } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Collapsible } from '@commonground/design-system'
import { IconServices } from '../../../../../../../icons'
import Amount from '../../../../../../../components/Amount'
import Table from '../../../../../../../components/Table'
import {
  DetailHeading,
  StyledCollapsibleBody,
} from '../../../../../../../components/DetailView'
import { OrganizationName, Separator } from './index.styles'

const Services = ({ services }) => {
  const { t } = useTranslation()
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
        {services.length ? (
          <Table role="grid" withLinks>
            <tbody>
              {services.map(({ service, organization }) => (
                <Table.Tr
                  to={`/directory/${organization}/${service}`}
                  key={`${organization}_${service}`}
                >
                  <Table.Td>
                    <OrganizationName>{organization}</OrganizationName>
                    <Separator> - </Separator>
                    <small>{service}</small>
                  </Table.Td>
                </Table.Tr>
              ))}
            </tbody>
          </Table>
        ) : (
          <small>{t('No services have been connected')}</small>
        )}
      </StyledCollapsibleBody>
    </Collapsible>
  )
}

Services.propTypes = {
  services: arrayOf(object).isRequired,
}

export default Services
