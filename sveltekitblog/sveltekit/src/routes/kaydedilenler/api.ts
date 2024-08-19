import { goto } from '$app/navigation';

export async function getSavedPosts(token: string) {
	const response = await fetch('http://localhost:3001/user/saved', {
		method: 'GET',
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json'
		}
	});

	if (!response.ok) {
		throw new Error('Kaydedilenler alınamadı. Status: ' + response.status);
	}
	return response.json();
}
