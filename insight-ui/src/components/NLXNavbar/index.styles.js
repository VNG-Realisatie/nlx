import styled from 'styled-components'
import Navigation from '../Navigation'

export const StyledNavbarNav = styled.nav`
  height: 56px;
  display: flex;
  align-items: center;
`

export const StyledNavbarLogoLink = styled.a`
  margin-right: 24px;
  flex: 0 0 50px;
`

export const StyledNavigation = styled(Navigation)`
  &:last-of-type {
    margin-right: auto;
  }
`
