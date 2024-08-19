/** @type {import('./$types').PageLoad} */
export async function getPostById(postID: number, token: string) {
	console.log('AAA');
	const response = await fetch(`http://localhost:3001/posts/${postID}`, {
		method: 'GET',
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json'
		}
	});

	if (!response.ok) {
		throw new Error('Yazı getirilemedi.');
	}

	return response.json();
}

export async function updatePost(
	postID: number,
	token: string,
	updatedPost: { postTitle: string; postDescription: string; tags: number[] }
) {
	const response = await fetch(`http://localhost:3001/posts/${postID}`, {
		method: 'PUT',
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(updatedPost)
	});

	if (!response.ok) {
		throw new Error('Yazı güncellenemedi.');
	}

	return response.json();
}
