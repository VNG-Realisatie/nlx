// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState, useEffect, createContext } from 'react'
import { string, func, node, shape } from 'prop-types'
import UserRepositoryOIDC from '../domain/user-repository-oidc'

const UserContext = createContext()

const UserContextProvider = ({
  children,
  fetchAuthenticatedUser,
  logout,
  user: defaultUser,
}) => {
  const [user, setUser] = useState(defaultUser || null)
  const [isReady, setReady] = useState(false)
  let canceled = false

  useEffect(
    () => {
      const fetchUser = async () => {
        try {
          const authenticatedUser = await fetchAuthenticatedUser()

          if (!canceled) {
            setUser(authenticatedUser)
          }
        } catch (err) {
          setUser(null)
        }

        setReady(true)
      }

      if (defaultUser) {
        setReady(true)
        return
      }

      fetchUser()
    },
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [],
  )

  const cancelFetchHandler = () => {
    canceled = true
  }

  const onLoginHandler = (user) => {
    setUser(user)
  }

  const onLogoutHandler = () => {
    setUser(null)
  }

  return (
    <UserContext.Provider
      value={{
        user: user,
        isReady: isReady,
        logout: logout,
        cancelFetch: cancelFetchHandler,
        onLoginHandler: onLoginHandler,
        onLogoutHandler: onLogoutHandler,
      }}
    >
      {children}
    </UserContext.Provider>
  )
}

UserContextProvider.propTypes = {
  fetchAuthenticatedUser: func,
  logout: func,
  children: node,
  user: shape({
    id: string,
    fullName: string,
    email: string,
    pictureUrl: string,
  }),
}

UserContextProvider.defaultProps = {
  fetchAuthenticatedUser: UserRepositoryOIDC.getAuthenticatedUser,
  logout: UserRepositoryOIDC.logout,
}

export default UserContext
export { UserContextProvider }
