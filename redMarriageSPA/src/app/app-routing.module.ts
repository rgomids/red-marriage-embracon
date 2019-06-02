import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { LandingPageComponent } from './components/landing-page/landing-page.component';
import { TimelineComponent } from './components/timeline/timeline.component';
import { PartnersComponent } from './components/partners/partners.component';

const routes: Routes = [
  {path: "land", component: LandingPageComponent},
  {path: "form", component: DashboardComponent},
  {path: "timeline", component: TimelineComponent},
  { path: "partners", component: PartnersComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
