import { Component } from '@angular/core';
import { bootstrapApplication } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { UserFormComponent } from './app/components/user-form/user-form.component';
import { importProvidersFrom } from '@angular/core';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [UserFormComponent],
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class App {}

bootstrapApplication(App, {
  providers: [
    importProvidersFrom(HttpClientModule)
  ]
});