/**
 * Copyright 2024 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { Component, EnvironmentInjector, Input, inject } from '@angular/core';
import { LOG_TOOL_ANNOTATOR_RESOLVER } from '../annotator/log-tool/resolver';
import { Subject, map, shareReplay, startWith, withLatestFrom } from 'rxjs';
import { InspectionDataStoreService } from '../services/inspection-data-store.service';
import { CommonModule } from '@angular/common';
import { LogHeaderComponent } from './header.component';
import { HighlightModule } from 'ngx-highlightjs';
import { ResolveTextPipe } from '../common/resolve-text.pipe';

@Component({
  selector: 'khi-log-body',
  templateUrl: './body.component.html',
  styleUrls: ['./body.component.sass'],
  imports:[
    CommonModule,
    LogHeaderComponent,
    HighlightModule,
    ResolveTextPipe
  ]
})
export class LogBodyComponent {
  private readonly dataStore = inject(InspectionDataStoreService);

  private readonly envInjector = inject(EnvironmentInjector);

  private readonly logToolAnnotatorResolver = inject(
    LOG_TOOL_ANNOTATOR_RESOLVER,
  );

  @Input()
  public set logIndex(index: number) {
    this.logIndexObservable.next(index);
  }

  private logIndexObservable = new Subject<number>();

  public logEntryObservable = this.logIndexObservable.pipe(
    startWith(0),
    withLatestFrom(this.dataStore.allLogs),
    map(([i, all]) => all[i]),
    shareReplay(1),
  );

  public logAnnotators = this.logToolAnnotatorResolver.getResolvedAnnotators(
    this.logEntryObservable,
    this.envInjector,
  );
}
