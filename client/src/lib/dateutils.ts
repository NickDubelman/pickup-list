export const monthNames = [
  'January',
  'February',
  'March',
  'April',
  'May',
  'June',
  'July',
  'August',
  'September',
  'October',
  'November',
  'December'
]

export const getStartOfPrevSunday = () => {
  const now = new Date()

  const curr = now
  let dayIndex = curr.getDay() // 0 for Sunday, 1 for Monday, etc..

  while (dayIndex > 0) {
    curr.setDate(curr.getDate() - 1)
    dayIndex--
  }

  curr.setHours(0, 0, 0, 0)
  return curr
}
