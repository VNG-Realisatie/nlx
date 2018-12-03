import React, {Fragment} from 'react'
import { Flex, Box } from '@rebass/grid'

import IconButton from 'src/IconButton/IconButton'
import FavoriteIcon from '@material-ui/icons/FavoriteBorder'

const FavoriteIconSmall = (<FavoriteIcon style={{fontSize: '14px'}}/>)
const FavoriteIconNormal = (<FavoriteIcon style={{fontSize: '20px'}}/>)
const FavoriteIconLarge = (<FavoriteIcon style={{fontSize: '24px'}}/>)

export const iconbuttonStory = (
    <Fragment>
        <Flex mb={3}>
            <Box mr={4}>
                <IconButton>{FavoriteIconNormal}</IconButton>
            </Box>
            <Box mr={4}>
                <IconButton disabled>{FavoriteIconNormal}</IconButton>
            </Box>
        </Flex>
        <Flex mb={3}>
            <Box mr={4}>
                <IconButton variant="secondary">{FavoriteIconNormal}</IconButton>
            </Box>
            <Box mr={4}>
                <IconButton variant="secondary" disabled>{FavoriteIconNormal}</IconButton>
            </Box>
        </Flex>
        <Flex mb={3}>
            <Box mr={4}>
                <IconButton variant="tertiary">{FavoriteIconNormal}</IconButton>
            </Box>
            <Box mr={4}>
                <IconButton variant="tertiary" disabled>{FavoriteIconNormal}</IconButton>
            </Box>
        </Flex>
        <Flex alignItems="center">
            <Box mr={4}>
                <IconButton size="small">{FavoriteIconSmall}</IconButton>
            </Box>
            <Box mr={4}>
                <IconButton size="normal">{FavoriteIconNormal}</IconButton>
            </Box>
            <Box mr={4}>
                <IconButton size="large">{FavoriteIconLarge}</IconButton>
            </Box>
        </Flex>
    </Fragment>
    )
