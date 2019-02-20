import styled from 'styled-components'
import Filters from '../../components/Filters/Filters'
import ServicesTableContainer from "../../containers/ServicesTableContainer/ServicesTableContainer";

export const StyledFilters = styled(Filters)`
  width: 600px;
  margin: 48px auto 32px;
`

export const StyledServicesTableContainer = styled(ServicesTableContainer)`
  margin-bottom: 56px;
`

StyledServicesTableContainer.displayName = 'ServicesTableContainer'
