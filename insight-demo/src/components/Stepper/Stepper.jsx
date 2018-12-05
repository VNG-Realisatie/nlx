import React, { PureComponent } from 'react'
import styled from 'styled-components'

const Wrapper = styled.div`
    position: relative;

    display: flex;
    justify-content: space-between;

    &:before {
        content: '';
        position: absolute;
        height: 3px;
        top: 14px;
        left: 16px;
        right: 16px;
        background-color: #EAEAEA;
    }
`

const Step = styled.div`
    position: relative;

    display: flex;
    align-items: center;
    justify-content: center;

    width: 32px;
    height: 32px;

    background-color: white;
    border-radius: 50%;
    border: 3px solid #EAEAEA;

    font-size: 16px;
    font-weight: bold;
    color: #999999;

    user-select: none;
    box-sizing: border-box;
`

class Stepper extends PureComponent {
    render() {
        return (
            <Wrapper>
                <Step>1</Step>
                <Step>2</Step>
                <Step>3</Step>
                <Step>4</Step>
            </Wrapper>
        )
    }
}

Stepper.propTypes = {
}

Stepper.defaultProps = {
}

export default Stepper
