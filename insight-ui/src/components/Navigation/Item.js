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
    color: #676D80;

    &:hover,
    &:focus {
      background-color: rgba(240, 242, 247, .5);
    }

    &:active {
      background-color: #F0F2F7;
    }

    &.active {
      background-color: #F1F5FF;
      color: #2961FF;
    }
  }
`
