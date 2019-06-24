import { css } from 'styled-components'

export function enterTransform(placement, animation) {
    let direction
    if (placement === 'top' || placement === 'bottom') {
        direction = 'Y'
    } else if (placement === 'left' || placement === 'right') {
        direction = 'X'
    }

    let distance = 10
    let effect = ''
    switch (animation) {
        case 'perspective':
            effect = `rotate${direction === 'Y' ? 'X' : 'Y'}(0)`
            break
        case 'scale':
            effect = 'scale(1)'
            break
        default:
            // no effect
    }

    const transform = `translate${direction}(${distance}px) ${effect}`
    return css`
        transform: ${transform};
    `
}

export function leaveTransform(placement, animation) {
    let direction
    if (placement === 'top' || placement === 'bottom') {
        direction = 'Y'
    } else if (placement === 'left' || placement === 'right') {
        direction = 'X'
    }

    let distance = 0
    let effect = ''
    switch (animation) {
        case 'perspective':
            const deg = placement === 'top' || placement === 'right' ? 60 : -60
            effect = `rotate${direction === 'Y' ? 'X' : 'Y'}(${deg}deg)`
            break
        case 'scale':
            effect = 'scale(0.5)'
            break
        case 'shift-toward':
            distance = placement === 'top' || placement === 'left' ? -20 : 20
            break
        default:
            // no effect
    }

    const transform = `translate${direction}(${distance}px) ${effect}`
    return css`
        transform: ${transform};
    `
}
