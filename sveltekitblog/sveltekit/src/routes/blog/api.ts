import { error } from '@sveltejs/kit';

export async function getCategories() {
	const response = await fetch('http://localhost:3001/categories');
	if (!response.ok) {
		throw error(response.status, 'Kategoriler alınamadı');
	}
	return await response.json();
}

export async function getPostsByCategory(categoryID: number) {
	const response = await fetch(`http://localhost:3001/posts?categoryID=${categoryID}`);
	if (!response.ok) {
		throw error(response.status, 'Yazılar alınamadı');
	}
	return await response.json();
}

export async function getAllPosts() {
	const response = await fetch('http://localhost:3001/posts');
	if (!response.ok) {
		throw new Error('Yazılar alınamadı');
	}
	return await response.json();
}

export async function savePost(postID: number, token: string) {
	const userID = localStorage.getItem('userID');

	const response = await fetch('http://localhost:3001/api/save', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${token}`
		},
		body: JSON.stringify({ postID: postID, userID: Number(userID) })
	});

	if (!response.ok) {
		throw new Error('Yazı kaydedilemedi');
	}

	return response.json();
}
