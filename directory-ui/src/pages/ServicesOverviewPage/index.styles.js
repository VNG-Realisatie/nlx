// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//
import styled from 'styled-components'
import Filters from '../../components/Filters'
import ServicesTableContainer from '../../containers/ServicesTableContainer'

export const StyledFilters = styled(Filters)`
  width: 100%;
  margin: 48px auto 32px;
`

export const StyledServicesTableContainer = styled(ServicesTableContainer)`
  margin-bottom: 56px;
`

StyledServicesTableContainer.displayName = 'ServicesTableContainer'
