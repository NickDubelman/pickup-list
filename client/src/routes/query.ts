import * as api from '$lib/api'

export async function post(req) {
  const resp = await api.post('query', req.body, req.locals.jwt)
  return { body: resp }
}
