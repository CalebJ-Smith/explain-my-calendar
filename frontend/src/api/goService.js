export async function ping() {
  try {
    const response = await fetch('/ping');
    return response.json().then((v) => v.message);
  } catch (error) {
    return []
  }
}