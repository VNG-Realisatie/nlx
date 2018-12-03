import React, {Fragment} from 'react'
import { Flex, Box } from '@rebass/grid'

import Button from 'src/Button/Button'
import FavoriteIcon from '@material-ui/icons/FavoriteBorder'

const icon = (
    <FavoriteIcon fontSize="small" />
)

export const buttonStory = (
    <Fragment>
        <Flex mb={3}>
            <Box mr={4}>
                <Button>Button</Button>
            </Box>
            <Box mr={4}>
                <Button icon={icon}>Button</Button>
            </Box>
            <Box mr={4}>
                <Button iconRight={icon}>Button</Button>
            </Box>
            <Box mr={4}>
                <Button disabled>Button</Button>
            </Box>
        </Flex>
        <Flex mb={3}>
            <Box mr={4}>
                <Button variant="secondary">Button</Button>
            </Box>
            <Box mr={4}>
                <Button variant="secondary" icon={icon}>Button</Button>
            </Box>
            <Box mr={4}>
                <Button variant="secondary" iconRight={icon}>Button</Button>
            </Box>
            <Box mr={4}>
                <Button variant="secondary" disabled>Button</Button>
            </Box>
        </Flex>
        <Flex mb={3}>
            <Box mr={4}>
                <Button variant="tertiary">Button</Button>
            </Box>
            <Box mr={4}>
                <Button variant="tertiary" icon={icon}>Button</Button>
            </Box>
            <Box mr={4}>
                <Button variant="tertiary" iconRight={icon}>Button</Button>
            </Box>
            <Box mr={4}>
                <Button variant="tertiary" disabled>Button</Button>
            </Box>
        </Flex>
        <Flex mb={3} alignItems="center">
            <Box mr={4}>
                <Button size="small">Button</Button>
            </Box>
            <Box mr={4}>
                <Button size="normal">Button</Button>
            </Box>
            <Box mr={4}>
                <Button size="large">Button</Button>
            </Box>
        </Flex>
    </Fragment>
)
