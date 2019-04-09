import styled from "styled-components";
import Card from "../../components/Card";

export const StyledContent = styled.div`
  flex: 1;
  background: #F7F9FC;
  display: flex;
  justify-content: center;
  align-items: center;
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
