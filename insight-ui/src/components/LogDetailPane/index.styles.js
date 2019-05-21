import styled from 'styled-components'

export const StyledLogDetailPane = styled.div`
  position: absolute;
  right: 0;
  top: 56px;
  width: 250px;
  background: #ffffff;
  box-shadow: 0 0 0 1px rgba(45,50,64,.05), 0 1px 8px rgba(45,50,64,.05);
  z-index: 1;
  min-height: calc(100% - 56px);
  padding: 20px 24px;
`

export const StyledTitle = styled.h3`
  color: #517FFF;
  font-size: 20px;
  font-weight: 700;
  line-height: 25px;
  overflow: hidden;
`

export const StyledCloseButton = styled.button`
  background: none;
  border: 0 none;
  padding: 5px;
  float: right;
  cursor: pointer;
`

export const StyledSubtitle = styled.h4`
  font-weight: 600;
  font-size: 12px;
  color: #2D3240;
  margin-bottom: 8px;
`

export const StyledDl = styled.dl`
  margin-top: 0;
  font-size: 12px;
  overflow: hidden;

  dt {
    color: #A3AABF;
    width: 75px;
    float: left;
    clear: both;
    padding-bottom: 8px;
  }

  dd {
    color: #2D3240;
    float: right;
    margin-left: 0;
    width: calc(100% - 75px);
    padding-bottom: 8px;
  }
`
