import styled from 'styled-components'
import Item from './Item'

const Navigation = styled.ul`
  display: flex;
  padding: 0;
  margin: 0;

  &:not(:last-of-type) {
    border-right: 1px solid #f0f2f7;
    padding-right: 10px;
    margin-right: 14px;
  }
`

Navigation.Item = Item

export default Navigation
