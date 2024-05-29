/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 *
 * For commercial use, a separate license must be obtained by purchasing from The Authors.
 * For commercial licensing inquiries, please contact The Authors listed in the AUTHORS file.
 */
import {
	BrowserWindow,
	GlobalShortcut,
} from 'electron';

export function registerZoomShortcuts(
	window: BrowserWindow,
	globalShortcut: GlobalShortcut
) {
	globalShortcut.register('CommandOrControl+Plus', () => {
		const currentZoom = window.webContents.getZoomFactor();
		window.webContents.setZoomFactor(currentZoom + 0.1);
	});

	globalShortcut.register('CommandOrControl+-', () => {
		const currentZoom = window.webContents.getZoomFactor();
		window.webContents.setZoomFactor(currentZoom - 0.1);
	});

	globalShortcut.register('CommandOrControl+0', () => {
		window.webContents.setZoomFactor(1.0);
	});
}