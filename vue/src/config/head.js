const defTitle = 'Hip Hop Blog'

var meta = {}

meta.title = (t) => {
  document.getElementById('meta-title').innerHTML = t + ' - ' + defTitle
}
meta.description = (d) => {
  document.getElementById('meta-desc').innerHTML = d
  document.getElementById('og-desc').innerHTML = d
}

export default meta
