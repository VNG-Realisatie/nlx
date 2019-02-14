import styled from 'styled-components'

export default styled.header`
  background: #ffffff;
  border-bottom: 1px solid #EAECF0;
  
  .container {
    width: 1140px;
    margin: 0 auto;
  }
  
  nav {
    padding: 2rem 0;
  }
  
  .navbar-logo {
    display: inline-block;
    vertical-align: middle;
    margin-right: 1.5rem;
  }
  
  .navbar-nav {
    display: inline-block;
    padding: 0;
    margin: 0;
    vertical-align: middle;
    
    &:not(:last-of-type) {
      border-right: 1px solid #F0F2F7;
      padding-right: 1rem;
    }
    
    .nav-item {
      display: inline-block;
      margin-left: 1.5rem;
      
      a {
        color: #A3AABF;
        font-size: 1rem;
        font-weight: 600;
        text-decoration: none;
      }
      
      &.active a {
        background: #F1F5FF;
        color: #517FFF;
        border-radius: 3px;
        padding: .3rem .7rem;
      }
    }
  }
  
  .nav-link {
    vertical-align: middle;
    display: inline-block;
    float: right;
  }
`
