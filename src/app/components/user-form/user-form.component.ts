import { Component, Input } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { UserService } from '../../services/user.service';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-user-form',
  standalone: true,
  imports: [FormsModule, CommonModule],
  templateUrl: './user-form.component.html',
  styleUrls: ['./user-form.component.css']
})
export class UserFormComponent {
  @Input() title = '';
  @Input() placeholder = '';
  
  userInput = '';
  message = '';
  messageColor = 'black';

  constructor(private userService: UserService) {}

  submitData() {
    if (!this.userInput.trim()) {
      this.message = 'Please enter data in format: NAME AGE CITY';
      this.messageColor = 'red';
      return;
    }

    this.userService.createUser(this.userInput).subscribe({
      next: (response) => {
        this.message = 'User data saved successfully!';
        this.messageColor = 'green';
        this.userInput = '';
      },
      error: (error) => {
        this.message = 'Error saving user data. Please try again.';
        this.messageColor = 'red';
      }
    });
  }
}