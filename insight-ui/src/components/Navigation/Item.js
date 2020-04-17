// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

export default styled.li`
  display: flex; /* Cancel li styles */
  margin-right: 4px;

  a {
    font-size: 14px;
    font-weight: 600;
    text-decoration: none;
    padding: 2px 10px 4px;
    border-radius: 3px;
    white-space: nowrap;
    color: #676d80;

    &:hover,
    &:focus {
      background-color: rgba(240, 242, 247, 0.5);
    }

    &:active {
      background-color: #f0f2f7;
    }

    &.active {
      background-color: #f1f5ff;
      color: #2961ff;
    }
  }
`
