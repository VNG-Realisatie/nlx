import styled from 'styled-components'
import Card from './components/Card'

export default styled.div`
  display: flex;
  height: 100vh;
`

export const StyledSidebar = styled.div`
  flex: 0 0 200px;
  background: #ffffff;
  box-shadow: 0 0 0 1px rgba(45,50,64,.05), 0 1px 8px rgba(45,50,64,.05);
  z-index: 1;
`

export const StyledSidebarHeader = styled.a`
  display: block;
  padding: 24px;
  color: #676D80;
  font-weight: 600;
  text-decoration: none;
  vertical-align: middle;
  
  svg {
    display: inline-block;
    width: 35px;
  }
`

export const StyledContent = styled.div`
  flex: 1;
  background: #F7F9FC;
  display: flex;
  justify-content: center;
  align-items: center;
`

export const StyledOrganizationList = styled.ul`
  list-style-type: none;
  padding: 0;
  
  li {
    &.active {
      color: #517FFF;
    
      a {
        border-left: 2px solid #517FFF;
      }
    }
  }
  
  a {
    color: #A3AABF;
    font-weight: 600;
    text-decoration: none;
    display: block;
    padding-left: 20px;
    line-height: 32px;
  }
`

export const StyledCard = styled(Card)`
  width: 400px;
  padding: 8px 24px 8px 24px;
  
  .text-muted {
    font-size: 12px;
    color: #A3AABF;
    
    a {
      text-decoration: none;
      color: #517FFF;
    }
  }
`
