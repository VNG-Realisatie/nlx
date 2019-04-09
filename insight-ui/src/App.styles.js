import styled from 'styled-components'
import { Navbar, Search } from '@commonground/design-system'

export default styled.div`
  display: flex;
  height: 100vh;
  flex-wrap: wrap;
  align-content: flex-start;
`

export const StyledNavbar = styled(Navbar)`
  flex: 1 100%;
  z-index: 3;
`

export const StyledContent = styled.div`
  flex: 1;
  background: #F7F9FC;
  display: flex;
  justify-content: center;
  align-items: center;
`
