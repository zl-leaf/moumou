function formatPath(parentPath: string, childPath: string): string {
  let fullPath = childPath
  if (parentPath.endsWith('/')) {
    parentPath = parentPath.slice(0, -1)
  }
  if (!fullPath.startsWith('/')) {
    fullPath = parentPath + '/' + fullPath
  }
  return fullPath
}

export default {
  formatPath
}