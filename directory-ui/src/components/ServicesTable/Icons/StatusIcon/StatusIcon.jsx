import styled from 'styled-components'

export default styled.div`
  display: inline-block;
  width: 10px;
  height: 10px;
  border-radius: 50%;
  border: 2px solid #63D19E;

  &[disabled] {
    border-color: #CAD0E0;
  }
`
