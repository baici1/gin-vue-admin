// 用于状态返回
var TagType = new Map()
TagType.set(0, 'danger')
TagType.set(1, 'success')
TagType.set(2, 'info')
TagType.set(3, 'warning')

export function ChooseTagType(t) {
  return TagType.get(t)
}
