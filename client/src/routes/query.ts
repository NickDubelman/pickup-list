import * as api from '$lib/api'

export async function post(req) {
  const resp = await api.post('query', req.body, req.locals.jwt)
  if (resp && resp.errors) {
    // FIXME: use error extensions
    if (resp.errors.map(e => e.message).includes('access token is expired')) {
      // attempt 'silent' refresh of JWT
      const body = {
        query: refreshMutation,
        variables: { input: { refreshToken: req.locals.refresh } }
      }

      const refreshResp = await api.post('query', body, req.locals.jwt)

      if (refreshResp && !refreshResp.errors) {
        const accessToken = refreshResp.data.refreshToken

        // retry original query with newly refreshed access token
        const resp = await api.post('query', req.body, accessToken)
        return {
          body: resp,
          headers: { 'set-cookie': `jwt=${accessToken}; Path=/; HttpOnly` }
        }
      }
    }
  }

  return { body: resp }
}

const refreshMutation = `
  mutation RefreshToken($input: RefreshTokenInput!){
    refreshToken(input: $input)
  }
`
