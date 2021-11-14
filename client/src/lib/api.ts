const base = 'http://localhost:8080'

async function send({ method, path, data, token }) {
  const opts = { method, headers: {}, body: undefined }

  if (data) {
    opts.headers['Content-Type'] = 'application/json'
    opts.body = JSON.stringify(data)
  }

  if (token) {
    opts.headers['Authorization'] = `Token ${token}`
  }

  return fetch(`${base}/${path}`, opts)
    .then(r => r.text())
    .then(json => {
      try {
        return JSON.parse(json)
      } catch (err) {
        return json
      }
    })
}

export function post(path, data, token) {
  return send({ method: 'POST', path, data, token })
}
