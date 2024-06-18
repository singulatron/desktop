import { Component } from '@angular/core';
import { UserService, Role, Permission } from '../../services/user.service';
import { first } from 'rxjs';

@Component({
	selector: 'app-roles',
	templateUrl: './roles.component.html',
	styleUrls: ['./roles.component.css'],
})
export class RolesComponent {
	roles: Role[] = [];
	permissions: Permission[] = [];
	selectedRole: Role | null = null;
	selectedRolePermissions: Set<string> = new Set<string>();

	roleSearchQuery: string = '';
	permissionSearchQuery: string = '';

	constructor(private userService: UserService) {
		this.userService.user$.pipe(first()).subscribe(() => {
			this.loggedInInit();
		});
	}

	async loggedInInit() {
		let rsp = await this.userService.getRoles();
		this.roles = await rsp.roles;

		let rsp2 = await this.userService.getPermissions();
		this.permissions = await rsp2.permissions;
	}

	selectRole(role: Role) {
		this.selectedRole = role;
		this.loadRolePermissions(role);
	}

	loadRolePermissions(role: Role) {
		this.selectedRolePermissions.clear();
		if (role.permissionIds) {
			role.permissionIds.forEach((id) => this.selectedRolePermissions.add(id));
		}
	}

	togglePermission(permissionId: string) {
		if (this.selectedRolePermissions.has(permissionId)) {
			this.selectedRolePermissions.delete(permissionId);
		} else {
			this.selectedRolePermissions.add(permissionId);
		}
	}

	async savePermissions() {
		if (this.selectedRole) {
			const permissionIds = Array.from(this.selectedRolePermissions);
			await this.userService.setRolePermissions(
				this.selectedRole.id as string,
				permissionIds
			);
			// Optionally, you can provide some user feedback here, such as a success message
		}
	}

	filteredRoles() {
		return this.roles.filter((role) =>
			role.name?.toLowerCase().includes(this.roleSearchQuery.toLowerCase())
		);
	}

	filteredPermissions() {
		return this.permissions.filter((permission) =>
			permission.name
				.toLowerCase()
				.includes(this.permissionSearchQuery.toLowerCase())
		);
	}
}
