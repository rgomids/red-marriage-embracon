import { BrowserModule } from '@angular/platform-browser';
//import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {BrowserAnimationsModule} from  '@angular/platform-browser/animations'
import {MatButtonModule} from '@angular/material/button';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatCardModule} from '@angular/material/card';
import {MatListModule} from '@angular/material/list';
import {MatMenuModule} from '@angular/material/menu';
import { NgModule } from '@angular/core';
import { MatStepperModule, MatFormFieldModule, MatInputModule } from '@angular/material';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { PerfilComponent } from './components/perfil/perfil.component';
import { TimelineComponent } from './components/timeline/timeline.component';
import { FornecedoresComponent } from './components/fornecedores/fornecedores.component';
import { HeaderComponent } from './components/header/header.component';
import { LandingPageComponent } from './components/landing-page/landing-page.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { ReactiveFormsModule } from '@angular/forms';



@NgModule({
  declarations: [
    AppComponent,
    PerfilComponent,
    TimelineComponent,
    FornecedoresComponent,
    HeaderComponent,
    LandingPageComponent,
    DashboardComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    AppRoutingModule,
    MatButtonModule, 
    MatCheckboxModule,
    MatGridListModule,
    MatCardModule,
    MatListModule,
    MatMenuModule,
    MatStepperModule,
    MatFormFieldModule,
    ReactiveFormsModule,
    MatInputModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
