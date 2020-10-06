// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import styled from 'styled-components'

export default styled.header`
  position: fixed;
  background: #ffffff;
  box-shadow: 0 0 0 1px rgba(45, 50, 64, 0.05), 0 1px 8px rgba(45, 50, 64, 0.05);
  z-index: 2;
  left: 0;
  right: 0;
  top: 0;

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
      border-right: 1px solid #f0f2f7;
      padding-right: 10px;
      margin-right: 14px;
    }

    .nav-item {
      display: flex; /* Cancel li styles */
      margin-right: 4px;

      a {
        font-size: 14px;
        font-weight: 600;
        text-decoration: none;
        padding: 2px 10px 4px;
        border-radius: 3px;
        white-space: nowrap;
      }

      &:not(.active) a {
        color: #a3aabf;

        &:hover,
        &:focus {
          background-color: #f7f9fc;
          color: #676d80;
        }

        &:active {
          background-color: #f0f2f7;
        }
      }

      &.active a {
        background-color: #f1f5ff;
        color: #517fff;
      }
    }
  }
`

export const StyledNavigation = styled.nav`
  height: 56px;
  display: flex;
  align-items: center;
`
