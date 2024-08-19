export async function getUserPosts(token: string) {
	const response = await fetch('http://localhost:3001/user/posts', {
		method: 'GET',
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json'
		}
	});

	if (!response.ok) {
		throw new Error('Yazılar alınamadı.');
	}

	return response.json();
}

export async function deletePost(postID: number, token: string) {
	const response = await fetch(`http://localhost:3001/posts/${postID}`, {
		method: 'DELETE',
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json'
		}
	});

	if (!response.ok) {
		throw new Error('Yazı silinemedi.');
	}
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
