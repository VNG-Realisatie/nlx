import styled from 'styled-components'
import Filters from '../../components/Filters/Filters'
import ServicesTableContainer from "../../containers/ServicesTableContainer/ServicesTableContainer";

export const StyledFilters = styled(Filters)`
  width: 600px;
  margin: 60px auto;
`

export const StyledServicesTableContainer = styled(ServicesTableContainer)`
  width: 1140px;
  margin: 0 auto 50px auto;
`.displayName = 'ServicesTableContainer'
