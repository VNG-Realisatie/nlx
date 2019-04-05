import styled from 'styled-components'

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
  
  svg {
    display: inline-block;
    width: 35px;
  }
`

export const StyledContent = styled.div`
  flex: 1;
  background: #F7F9FC;
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
