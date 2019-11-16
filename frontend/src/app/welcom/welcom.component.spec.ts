import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { WelcomComponent } from './welcom.component';
import { LayoutHeaderSidebarContentsComponent } from '../layout/layout-header-sidebar-contents/layout-header-sidebar-contents.component';
import { HeaderComponent } from '../header/header.component';

describe('WelcomComponent', () => {
  let component: WelcomComponent;
  let fixture: ComponentFixture<WelcomComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        WelcomComponent,
        LayoutHeaderSidebarContentsComponent,
        HeaderComponent,
      ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(WelcomComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
