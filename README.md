# GoToAzure
This is a sample project to provide a quick-start for Golang developers to experience the deployment to their Golang applications to Microsoft Azure Web Apps.

## Section 1: Setup Build Pipeline
1. Login to https://dev.azure.com;
2. Click on the "+ New Project" button at the top-right corner (It may not be visible to you if you are not admin);
3. Click "Create" button to proceed after keying in the information about your project and deciding whether your repo should be Public or Private;
4. Choose "Builds" under Pipelines at the left menu;
5. Click on the "New Pipeline" button;
   ![](github-images/builds-new-pipeline.png?raw=true)
6. Click on the link which says "Use the classic editor";
   ![](github-images/use-classic-editor.png?raw=true)
7. Select the location of your repo (eg. Azure DevOps or GitHub, in this readme, I use GitHub because the code is on GitHub);
8. Continue to next step after connecting to the repo;
9. Search "go" in the search box and choose "Go (Preview)" as template;
   ![](github-images/golang-template.png?raw=true)
10. Leave the settings of the Pipeline as they are;
11. Visit the "Triggers" tab;
12. Enable the Continuous Integration and batch processing;
   ![](github-images/enable-continuous-integration.png?raw=true)
13. Click "Save & Queue" button to proceed;
   ![](github-images/save-and-queue.png?raw=true)
14. Wait for the Build process to finish with all green ticks;
15. Celebrate for setting up the Build Pipeline when the Artifact file is successfully generated and published.
   ![](github-images/artifacts-drop.png?raw=true)

## Section 2: Setup Azure Web App
1. Login to https://portal.azure.com;
2. Create a new Azure web app with the following options configured accordingly;
   - Publish: Code
   - Runtime Stack: ASP .NET v4.7
   - Operating System: Windows
   ![](github-images/create-app-service.png?raw=true)
3. Wait for the message saying "Your deployment is complete";
   ![](github-images/app-service-deployment-done.png?raw=true)
4. Celebrate for setting up new Azure web app successfully.

## Section 3: Setup Release Pipeline
1. Back to Azure DevOps (https://dev.azure.com);
2. Choose the project created in Section 1;
3. Choose "Releases" under Pipelines at the left menu;
4. Click on the "New Pipeline" button;
5. Apply on the "Azure App Service deployment" template;
   ![](github-images/apply-release-template.png?raw=true)
6. Click on the "Add an artifact" box in the diagram and choose the Build pipeline created in Section 1;
   ![](github-images/choose-artifact.png?raw=true)
7. Click "Add" button to proceed;
8. Enable the Continuous Deployment trigger;
   ![](github-images/continuous-deployment-trigger.png?raw=true)
9. Edit the Stage 1 by clicking on the "1 job, 1 task";
10. Connect with Azure by choosing the correct corresponding Azure Subscription and App Service Name for the web app created in Section 2 (You may need to click on the Authorize button to authorize this action so that Azure service connection can be configured);
   ![](github-images/linked-to-azure.png?raw=true)
11. Click on the "+" button beside "Run on agent";
12. Add two "Azure App Service Manage" tasks and surround the "Deploy Azure App Service" task, as shown in the screenshot below;
   ![](github-images/add-azure-app-service-manage-tasks.png?raw=true)
13. Configured the two newly added tasks so that the first is to stop the Web App and the second is to start the Web App;
   ![](github-images/stop-and-start-webapp.png?raw=true)
14. Click on the "Deploy Azure App Service" task;
15. Add in the following line under "Generate web.config parameters for Python, Node.js, Go and Java apps";
    > -GoExeFileName "$(System.TeamProject).exe" -appType Go
   ![](github-images/webconfig-for-golang.png?raw=true)
16. Tick the checkbox "Select deployment method" and make sure both "Web Deploy" and "Take App Offline" is selected;
   ![](github-images/select-deployment-method.png?raw=true)
17. Save the pipeline;
18. Click "Create release";
19. Choose "Releases" under Pipelines at the left menu;
20. Wait for the "Release-1" to be having green tick;
   ![](github-images/released.png?raw=true)
21. Visit the web app to check if the deployment is successful;
22. Celebrate if you see the following page (You can also test /api to see JSON object returned).
   ![](github-images/success.png?raw=true)

## Important Notes
The reason why we need to have two tasks in the Release Pipeline to first stop the Web App after the new code is deployed is that if we do not stop the Web App first then the code cannot be deployed with the following error message.

> Error: Error Code: ERROR_FILE_IN_USE
> 
> More Information: Web Deploy cannot modify the file GoToAzure.exe on the destination because it is locked by an external process. In order to allow the publish operation to succeed, you may need to either restart your application to release the lock, or use the AppOffline rule handler for .Net applications on your next publish attempt. Learn more at: http://go.microsoft.com/fwlink/?LinkId=221672#ERROR_FILE_IN_USE.
> 
> Error count: 1

In the app.go, we have the getPort() function is to get the actual port used in Azure Web App for the program. Hence, please do not hardcode ":80", instead read the port number from the HTTP_PLATFORM_PORT.

```
func getPort() string {
    p := os.Getenv("HTTP_PLATFORM_PORT")
    if p != "" {
	      return ":" + p   
    }
    return ":80"
}
```

## Contributions are Welcome!

This project will continue to evolve. We will do modifications all the time as long as Azure DevOps and Azure Portal do not stop changing.

Any help will be appreciated. You can develop new features, fix bugs, improve the documentation, or do some other cool stuff.

If you have new ideas or want to complain about bugs, feel free to [create a new issue](https://github.com/goh-chunlin/GoToAzure/issues/new).

Let's build the best documentation for [Golang with Azure](https://medium.com/golang-with-azure)!

## Code of Conduct

This project has adopted the code of conduct defined by the [Contributor Covenant](http://contributor-covenant.org/)
to clarify expected behavior in our community.

## Azure Community Singapore (ACS)

This project is supported by the [Azure Community Singapore (ACS)](https://www.meetup.com/AzureSG/).

## References
- [Deploy Golang App to Azure Web Apps with CI/CD on DevOps](https://medium.com/golang-with-azure/after-we-have-our-code-on-github-repository-now-its-time-to-automate-our-builds-and-deployments-2e332790f3)
