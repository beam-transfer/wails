// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

/**
 * GreetService is great
 * @module
 */

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Call as $Call, Create as $Create} from "/wails/runtime.js";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as services$0 from "./services/models.js";

/**
 * Greet does XYZ
 */
export function Greet(name: string, title: services$0.Title): Promise<string> & { cancel(): void } {
    let $resultPromise = $Call.ByName("main.GreetService.Greet", name, title) as any;
    return $resultPromise;
}
