export async function getPostById(postID: number, token: string) {
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

export async function getCommentsByPostId(postID: number) {
	const response = await fetch(`http://localhost:3001/comments?postID=${postID}`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json'
		}
	});

	if (!response.ok) {
		throw new Error('Yorumlar getirilemedi.');
	}

	return response.json();
}

export async function createComment(postID: number, commentt: string, token: string) {
	const response = await fetch(`http://localhost:3001/comments`, {
		method: 'POST',
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({ postID, commentt })
	});

	if (!response.ok) {
		const errorData = await response.json();
		throw new Error(`Yorum oluşturulamadı: ${errorData.message}`);
	}

	return response.json();
}
