import styled from 'styled-components'

export default styled.header`
  padding: 0 36px;
  background: #ffffff;
  border-bottom: 1px solid #EAECF0;

  .navbar-logo {
    margin-right: 24px;

    img {
      height: 16px;
    }
  }

  .navbar-gitlab {
    margin-left: auto;

    img {
      height: 20px;
    }
  }

  .navbar-nav {
    display: flex;
    padding: 0;
    margin: 0;

    &:not(:last-of-type) {
      border-right: 1px solid #F0F2F7;
      padding-right: 10px;
      margin-right: 14px;
    }

    .nav-item {
      display: flex; /* Cancel li styles */
      margin-right: 4px;

      a {
        font-size: 14px;
        text-decoration: none;
        padding: 4px 10px 6px;
        border-radius: 3px;
      }

      &:not(.active) a {
        color: #A3AABF;

        &:hover {
          background-color: #F7F9FC;
        }
      }

      &.active a {
        background-color: #F1F5FF;
        color: #517FFF;
      }
    }
  }
`

export const StyledNavigation = styled.nav`
  height: 56px;
  display: flex;
  align-items: center;
`