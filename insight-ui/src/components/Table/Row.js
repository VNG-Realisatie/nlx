import styled from 'styled-components'

const Row = styled.tr`
  display: table-row;

  &:first-child {
    td {
      &:first-child {
        border-top-left-radius: 3px;
      }

      &:last-child {
        border-top-right-radius: 3px;
      }
    }
  }

  &:not(:last-child) {
    td {
      border-bottom: 1px solid #f0f2f7;
    }
  }

  &:last-child {
    td {
      &:first-child {
        border-bottom-left-radius: 3px;
      }

      &:last-child {
        border-bottom-right-radius: 3px;
      }
    }
  }
`

export default Row
