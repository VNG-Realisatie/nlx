import styled, { keyframes } from 'styled-components'

const topValues = [50, 54, 57, 58, 57, 54, 50, 45]
const leftValues = [50, 45, 39, 32, 25, 19, 14, 10]

const generateAnimations = () =>
  Array.from({ length: 8 }).map(
    (value, i) => `
      &:nth-child(${i + 1}) {
        animation-delay: ${-0.036 * (i + 1)}s;
      }
  
      &:nth-child(${i + 1}):after {
        top: ${topValues[i]}px;
        left: ${leftValues[i]}px;
      }
    `,
  )

const rotate = keyframes`
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
`

export const StyledSpinner = styled.div`
  transform: translate(-50%, -50%);
  position: absolute;
  top: 50%;
  left: 50%;
`

export const StyledBulletContainer = styled.div`
  display: inline-block;
  position: relative;
  width: 64px;
  height: 64px;
`

export const StyledBullet = styled.div`
  animation: 1.2s ${rotate} cubic-bezier(0.5, 0, 0.5, 1) infinite;
  transform-origin: 32px 32px;
  background-color: #febf24;

  &:after {
    background: #febf24;
    border-radius: 50%;
    content: ' ';
    display: block;
    margin: -3px 0 0 -3px;
    position: absolute;
    width: 6px;
    height: 6px;
  }

  ${generateAnimations()}
`
