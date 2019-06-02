import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { LandingPageComponent } from './components/landing-page/landing-page.component';
import { TimelineComponent } from './components/timeline/timeline.component';

const routes: Routes = [
  {path: "land", component: LandingPageComponent},
  {path: "form", component: DashboardComponent},
  {path: "timeline", component: TimelineComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
